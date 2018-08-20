package handlers

import (
	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
	"github.com/taeho-io/family/services/todo_groups/repos/todo_groups_repo"
	"golang.org/x/net/context"
)

type DeleteTodoGroupFunc func(
	ctx context.Context,
	req *todo_groups.DeleteTodoGroupRequest,
) (*todo_groups.DeleteTodoGroupResponse, error)

func DeleteTodoGroup(todoGroupsTable *todo_groups_repo.Table) DeleteTodoGroupFunc {
	return func(
		ctx context.Context,
		req *todo_groups.DeleteTodoGroupRequest,
	) (*todo_groups.DeleteTodoGroupResponse, error) {
		err := todoGroupsTable.DeleteByID(req.GetTodoGroupId())
		if err != nil {
			return nil, err
		}

		return &todo_groups.DeleteTodoGroupResponse{}, nil
	}
}
