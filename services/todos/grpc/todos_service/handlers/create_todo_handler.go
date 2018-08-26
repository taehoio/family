package handlers

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
	"github.com/taeho-io/family/services/base/grpc/base_service"
	"github.com/taeho-io/family/services/todos/models"
	"github.com/taeho-io/family/services/todos/repos/todos_repo"
)

type CreateTodoFunc func(ctx context.Context, req *todos.CreateTodoRequest) (*todos.CreateTodoResponse, error)

func CreateTodo(
	todosTable *todos_repo.Table,
	hasPermissionByAccountID base_service.HasPermissionByAccountIDFunc,
) CreateTodoFunc {
	return func(ctx context.Context, req *todos.CreateTodoRequest) (*todos.CreateTodoResponse, error) {
		if err := validateTodoInput(req); err != nil {
			return nil, err
		}

		if err := hasPermissionByAccountID(ctx, req.AccountId); err != nil {
			return nil, err
		}

		logger := ctxlogrus.Extract(ctx).WithField("req", req)

		todo := models.NewTodoFromProto(req.Todo)
		todo.TodoID = xid.New().String()
		if err := todosTable.Put(todo); err != nil {
			logger.WithFields(logrus.Fields{
				"what": "todosTable.Put",
				"todo": todo,
			}).Error(err)

			return nil, err
		}

		return &todos.CreateTodoResponse{
			Todo: todo.ToProto(),
		}, nil
	}
}

func validateTodoInput(req *todos.CreateTodoRequest) error {
	if req.Todo.Title == "" {
		return InvalidTitleError
	}
	if req.Todo.ParentId == "" {
		return InvalidParentIDError
	}

	return nil
}
