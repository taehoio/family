package handlers

import (
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
	"github.com/taeho-io/family/services/todos/repos/todos_repo"
)

type CreateTodoFunc func(ctx context.Context, req *todos.CreateTodoRequest) (*todos.CreateTodoResponse, error)

func CreateTodo(todosTable *todos_repo.Table) CreateTodoFunc {
	return func(ctx context.Context, req *todos.CreateTodoRequest) (*todos.CreateTodoResponse, error) {
		return nil, nil
	}
}
