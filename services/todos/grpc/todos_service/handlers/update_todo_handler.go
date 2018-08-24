package handlers

import (
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
	"github.com/taeho-io/family/services/todos/repos/todos_repo"
)

type UpdateTodoFunc func(ctx context.Context, req *todos.UpdateTodoRequest) (*todos.UpdateTodoResponse, error)

func UpdateTodo(todosTable *todos_repo.Table) UpdateTodoFunc {
	return func(ctx context.Context, req *todos.UpdateTodoRequest) (*todos.UpdateTodoResponse, error) {
		return nil, nil
	}
}
