package handlers

import (
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
	"github.com/taeho-io/family/services/todo_groups/models"
	"github.com/taeho-io/family/services/todo_groups/repos/todo_groups_repo"
)

type CreateTodoGroupFunc func(
	ctx context.Context, req *todo_groups.CreateTodoGroupRequest,
) (*todo_groups.CreateTodoGroupResponse, error)

func CreateTodoGroup(todoGroupTable *todo_groups_repo.Table) CreateTodoGroupFunc {
	return func(
		ctx context.Context,
		req *todo_groups.CreateTodoGroupRequest,
	) (*todo_groups.CreateTodoGroupResponse, error) {
		err := validateTodoGroupInput(req)
		if err != nil {
			return nil, err
		}

		todoGroup := models.NewTodoGroupFromProto(req.GetTodoGroup())
		todoGroup.TodoGroupID = uuid.NewV4().String()
		if err := todoGroupTable.Put(todoGroup); err != nil {
			return nil, err
		}

		return &todo_groups.CreateTodoGroupResponse{
			TodoGroup: todoGroup.ToProto(),
		}, nil
	}
}

func validateTodoGroupInput(req *todo_groups.CreateTodoGroupRequest) error {
	if req.TodoGroup.Title == "" {
		return InvalidTitleError
	}

	return nil
}
