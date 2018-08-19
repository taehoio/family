package handlers

import (
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
	"github.com/taeho-io/family/services/todo_groups/repos/todo_groups_repo"
)

type GetTodoGroupFunc func(
	ctx context.Context,
	req *todo_groups.GetTodoGroupRequest,
) (*todo_groups.GetTodoGroupResponse, error)

func GetTodoGroup(todoGroupTable *todo_groups_repo.Table) GetTodoGroupFunc {
	return func(
		ctx context.Context,
		req *todo_groups.GetTodoGroupRequest,
	) (*todo_groups.GetTodoGroupResponse, error) {
		todoGroup, err := todoGroupTable.GetByID(req.GetTodoGroupId())
		if err != nil {
			return nil, err
		}

		return &todo_groups.GetTodoGroupResponse{
			TodoGroup: todoGroup.ToProto(),
		}, nil
	}
}
