package handler

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/todos/internal/repo"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todogroups"
	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
)

type DeleteTodoFunc func(ctx context.Context, req *todos.DeleteTodoRequest) (*todos.DeleteTodoResponse, error)

func DeleteTodo(
	todosRepo repo.TodosRepo,
	getAccountIDFromContext base.GetAccountIDFromContextFunc,
	hasPermissionByAccountID base.HasPermissionByAccountIDFunc,
	todoGroupsServiceClient todogroups.TodoGroupsServiceClient,
) DeleteTodoFunc {
	return func(ctx context.Context, req *todos.DeleteTodoRequest) (*todos.DeleteTodoResponse, error) {
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

		permitType := todoGroupRes.TodoGroup.PermitType
		if !(permitType == todogroups.PermitType_PERMIT_TYPE_OWNER || permitType == todogroups.PermitType_PERMIT_TYPE_EDITOR) {
			return nil, base.PermissionDeniedError
		}

		if err := todosRepo.DeleteByID(req.TodoId); err != nil {
			logger.WithFields(logrus.Fields{
				"what":   "todosRepo.DeleteByID",
				"todoId": req.TodoId,
			}).Error(err)

			return nil, err
		}

		return &todos.DeleteTodoResponse{}, nil
	}
}
