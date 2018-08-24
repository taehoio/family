package handlers

import (
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
	"github.com/taeho-io/family/services/todos/repos/todos_repo"
)

type ListTodosFunc func(ctx context.Context, req *todos.ListTodosRequest) (*todos.ListTodosResponse, error)

func ListTodos(todosTable *todos_repo.Table) ListTodosFunc {
	return func(ctx context.Context, req *todos.ListTodosRequest) (*todos.ListTodosResponse, error) {
		return nil, nil
	}
}
