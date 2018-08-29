package handler

import (
	"golang.org/x/net/context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todogroups"
	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/todos/internal/repo"
)

type GetTodoFunc func(ctx context.Context, req *todos.GetTodoRequest) (*todos.GetTodoResponse, error)

func GetTodo(
	todosRepo repo.TodosRepo,
	getAccountIDFromContext base.GetAccountIDFromContextFunc,
	hasPermissionByAccountID base.HasPermissionByAccountIDFunc,
	todoGroupsServiceClient todogroups.TodoGroupsServiceClient,
) GetTodoFunc {
	return func(ctx context.Context, req *todos.GetTodoRequest) (*todos.GetTodoResponse, error) {
		req.AccountId = getAccountIDFromContext(ctx)
		if err := hasPermissionByAccountID(ctx, req.AccountId); err != nil {
			return nil, err
		}

		logger := ctxlogrus.Extract(ctx).WithField("req", req)

		todo, err := todosRepo.GetByID(req.TodoId)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"what":   "todosRepo.GetByID",
				"todoId": "req.TodoId",
			})
			return nil, err
		}

		if todo.ParentType != todos.ParentType(todos.ParentType_PARENT_TYPE_TODO_GROUP) {
			return nil, InvalidParentTypeError
		}

		getTogoGroupReq := &todogroups.GetTodoGroupRequest{
			AccountId:   req.AccountId,
			TodoGroupId: todo.ParentID,
		}
		todoGroupRes, err := todoGroupsServiceClient.GetTodoGroup(ctx, getTogoGroupReq)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"what":            "todoGroupsServiceClient.GetTodoGroup",
				"getTogoGroupReq": getTogoGroupReq,
			}).Error(err)

			return nil, err
		}
		if todoGroupRes.TodoGroup == nil {
			return nil, base.PermissionDeniedError
		}

		return &todos.GetTodoResponse{
			Todo: todo.ToProto(),
		}, nil
	}
}
