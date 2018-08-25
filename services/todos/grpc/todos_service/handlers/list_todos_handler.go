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

type ListTodosFunc func(ctx context.Context, req *todos.ListTodosRequest) (*todos.ListTodosResponse, error)

func ListTodos(
	todosTable *todos_repo.Table,
	getAccountIDFromContext base_service.GetAccountIDFromContextFunc,
	hasPermissionByAccountID base_service.HasPermissionByAccountIDFunc,
	todoGroupsServiceClient todo_groups.TodoGroupsServiceClient,
) ListTodosFunc {
	return func(ctx context.Context, req *todos.ListTodosRequest) (*todos.ListTodosResponse, error) {
		req.AccountId = getAccountIDFromContext(ctx)
		if err := hasPermissionByAccountID(ctx, req.AccountId); err != nil {
			return nil, err
		}

		logger := ctxlogrus.Extract(ctx)

		if req.ParentType != todos.ParentType_PARENT_TYPE_TODO_GROUP {
			return nil, InvalidParentTypeError
		}

		getTogoGroupReq := &todo_groups.GetTodoGroupRequest{
			AccountId:   req.AccountId,
			TodoGroupId: req.ParentId,
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

		todoList, err := todosTable.ListByParentID(req.ParentId)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"what":     "todoTable.ListByParendID",
				"parentId": req.ParentId,
			}).Error(err)

			return nil, err
		}

		var todoProtoList []*todos.Todo
		for _, todo := range todoList {
			todoProtoList = append(todoProtoList, todo.ToProto())
		}

		return &todos.ListTodosResponse{
			Todos: todoProtoList,
		}, nil
	}
}
