package handlers

import (
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
	"github.com/taeho-io/family/services/todos/repos/todos_repo"
)

type GetTodoFunc func(ctx context.Context, req *todos.GetTodoRequest) (*todos.GetTodoResponse, error)

func GetTodo(todosTable *todos_repo.Table) GetTodoFunc {
	return func(ctx context.Context, req *todos.GetTodoRequest) (*todos.GetTodoResponse, error) {
		return nil, nil
	}
}
