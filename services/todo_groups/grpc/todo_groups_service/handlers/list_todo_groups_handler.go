package handlers

import (
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
	"github.com/taeho-io/family/services/todo_groups/repos/todo_groups_repo"
)

type ListTodoGroupsFunc func(
	ctx context.Context,
	req *todo_groups.ListTodoGroupsRequest,
) (*todo_groups.ListTodoGroupsResponse, error)

func ListTodoGroups(todoGroupTable *todo_groups_repo.Table) ListTodoGroupsFunc {
	return func(
		ctx context.Context,
		req *todo_groups.ListTodoGroupsRequest,
	) (*todo_groups.ListTodoGroupsResponse, error) {
		todoGroups, err := todoGroupTable.ListByIDs([]string{})
		if err != nil {
			return nil, err
		}

		var todoGroupProtos []*todo_groups.TodoGroup
		for _, todoGroup := range todoGroups {
			todoGroupProtos = append(todoGroupProtos, todoGroup.ToProto())
		}

		return &todo_groups.ListTodoGroupsResponse{
			TodoGroups: todoGroupProtos,
		}, nil
	}
}
