package handlers

import (
	"fmt"

	"github.com/satori/go.uuid"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
	"github.com/taeho-io/family/services/todos/models"
	"github.com/taeho-io/family/services/todos/repos/todo_groups_repo"
)

var (
	InvalidTitleError = fmt.Errorf("invalid title")
)

type CreateTodoGroupFunc func(ctx context.Context, req *todos.CreateTodoGroupRequest) (*todos.CreateTodoGroupResponse, error)

func CreateTodoGroup(todoGroupTable *todo_groups_repo.Table) CreateTodoGroupFunc {
	return func(ctx context.Context, req *todos.CreateTodoGroupRequest) (*todos.CreateTodoGroupResponse, error) {
		err := validateTodoGroupInput(req)
		if err != nil {
			return nil, err
		}

		todoGroup := models.NewTodoGroupFromProto(req.TodoGroup)
		todoGroup.TodoGroupID = uuid.NewV4().String()
		if err := todoGroupTable.Put(todoGroup); err != nil {
			return nil, err
		}

		return &todos.CreateTodoGroupResponse{
			TodoGroup: todoGroup.ToProto(),
		}, nil
	}
}

func validateTodoGroupInput(req *todos.CreateTodoGroupRequest) error {
	if req.TodoGroup.Title == "" {
		return InvalidTitleError
	}

	return nil
}
