package handler

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/todos/internal/model"
	"github.com/taeho-io/family/services/todos/internal/repo"
)

type CreateTodoFunc func(ctx context.Context, req *todos.CreateTodoRequest) (*todos.CreateTodoResponse, error)

func CreateTodo(
	todosRepo repo.TodosRepo,
	hasPermissionByAccountID base.HasPermissionByAccountIDFunc,
) CreateTodoFunc {
	return func(ctx context.Context, req *todos.CreateTodoRequest) (*todos.CreateTodoResponse, error) {
		if err := validateTodoInput(req); err != nil {
			return nil, err
		}

		if err := hasPermissionByAccountID(ctx, req.AccountId); err != nil {
			return nil, err
		}

		logger := ctxlogrus.Extract(ctx).WithField("req", req)

		todo := model.NewTodoFromProto(req.Todo)
		todo.TodoID = xid.New().String()
		if err := todosRepo.Put(todo); err != nil {
			logger.WithFields(logrus.Fields{
				"what": "todosRepo.Put",
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
