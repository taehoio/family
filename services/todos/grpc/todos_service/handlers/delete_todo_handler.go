package handlers

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
	"github.com/taeho-io/family/services/base/grpc/base_service"
	"github.com/taeho-io/family/services/todos/repos/todos_repo"
)

type DeleteTodoFunc func(ctx context.Context, req *todos.DeleteTodoRequest) (*todos.DeleteTodoResponse, error)

func DeleteTodo(
	todosTable todos_repo.IFace,
	getAccountIDFromContext base_service.GetAccountIDFromContextFunc,
	hasPermissionByAccountID base_service.HasPermissionByAccountIDFunc,
	todoGroupsServiceClient todo_groups.TodoGroupsServiceClient,
) DeleteTodoFunc {
	return func(ctx context.Context, req *todos.DeleteTodoRequest) (*todos.DeleteTodoResponse, error) {
		req.AccountId = getAccountIDFromContext(ctx)
		if err := hasPermissionByAccountID(ctx, req.AccountId); err != nil {
			return nil, err
		}

		logger := ctxlogrus.Extract(ctx).WithField("req", req)

		todo, err := todosTable.GetByID(req.TodoId)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"what":   "todosTable.GetByID",
				"todoId": "req.TodoId",
			})
			return nil, err
		}

		if todo.ParentType != todos.ParentType(todos.ParentType_PARENT_TYPE_TODO_GROUP) {
			return nil, InvalidParentTypeError
		}

		getTogoGroupReq := &todo_groups.GetTodoGroupRequest{
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
			return nil, base_service.PermissionDeniedError
		}

		permitType := todoGroupRes.TodoGroup.PermitType
		if !(permitType == todo_groups.PermitType_PERMIT_TYPE_OWNER || permitType == todo_groups.PermitType_PERMIT_TYPE_EDITOR) {
			return nil, base_service.PermissionDeniedError
		}

		if err := todosTable.DeleteByID(req.TodoId); err != nil {
			logger.WithFields(logrus.Fields{
				"what":   "todosTable.DeleteByID",
				"todoId": req.TodoId,
			}).Error(err)

			return nil, err
		}

		return &todos.DeleteTodoResponse{}, nil
	}
}
