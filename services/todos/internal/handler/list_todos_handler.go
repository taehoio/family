package handler

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todogroups"
	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/todos/internal/repo"
)

type ListTodosFunc func(ctx context.Context, req *todos.ListTodosRequest) (*todos.ListTodosResponse, error)

func ListTodos(
	todosRepo repo.TodosRepo,
	todoGroupsServiceClient todogroups.TodoGroupsServiceClient,
) ListTodosFunc {
	return func(ctx context.Context, req *todos.ListTodosRequest) (*todos.ListTodosResponse, error) {
		accountID := base.GetAccountIDFromContext(ctx)
		if err := base.HasPermissionByAccountID(ctx, accountID); err != nil {
			return nil, err
		}

		logger := ctxlogrus.Extract(ctx).WithField("req", req)

		if req.ParentType != todos.ParentType_PARENT_TYPE_TODO_GROUP {
			return nil, InvalidParentTypeError
		}

		getTogoGroupReq := &todogroups.GetTodoGroupRequest{
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
			return nil, base.PermissionDeniedError
		}

		todoList, err := todosRepo.ListByParentID(req.ParentId)
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
