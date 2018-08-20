package handlers

import (
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
	"github.com/taeho-io/family/services/todo_groups/repos/todo_groups_repo"
)

type UpdateTodoGroupFunc func(
	ctx context.Context, req *todo_groups.UpdateTodoGroupRequest,
) (*todo_groups.UpdateTodoGroupResponse, error)

func UpdateTodoGroup(todoGroupsTable *todo_groups_repo.Table) UpdateTodoGroupFunc {
	return func(
		ctx context.Context,
		req *todo_groups.UpdateTodoGroupRequest,
	) (*todo_groups.UpdateTodoGroupResponse, error) {
		_, err := todoGroupsTable.UpdateTitle(
			req.GetTodoGroupId(),
			req.GetTodoGroup().GetTitle(),
		)
		if err != nil {
			return nil, err
		}

		todoGroup, err := todoGroupsTable.UpdateDescription(
			req.GetTodoGroupId(),
			req.GetTodoGroup().GetDescription(),
		)
		if err != nil {
			return nil, err
		}

		return &todo_groups.UpdateTodoGroupResponse{
			TodoGroup: todoGroup.ToProto(),
		}, nil
	}
}
