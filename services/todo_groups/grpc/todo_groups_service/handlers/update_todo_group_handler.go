package handlers

import (
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
	"github.com/taeho-io/family/services/todo_groups/repos/todo_groups_repo"
)

type UpdateTodoGroupFunc func(
	ctx context.Context, req *todo_groups.UpdateTodoGroupRequest,
) (*todo_groups.UpdateTodoGroupResponse, error)

func UpdateTodoGroup(_ *todo_groups_repo.Table) UpdateTodoGroupFunc {
	return func(
		ctx context.Context,
		req *todo_groups.UpdateTodoGroupRequest,
	) (*todo_groups.UpdateTodoGroupResponse, error) {
		return &todo_groups.UpdateTodoGroupResponse{
			TodoGroup: req.GetTodoGroup(),
		}, nil
	}
}
