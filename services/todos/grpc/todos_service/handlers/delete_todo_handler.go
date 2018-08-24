package handlers

import (
	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
	"github.com/taeho-io/family/services/todos/repos/todos_repo"
	"golang.org/x/net/context"
)

type DeleteTodoFunc func(ctx context.Context, req *todos.DeleteTodoRequest) (*todos.DeleteTodoResponse, error)

func DeleteTodo(todosTable *todos_repo.Table) DeleteTodoFunc {
	return func(ctx context.Context, req *todos.DeleteTodoRequest) (*todos.DeleteTodoResponse, error) {
		return nil, nil
	}
}
