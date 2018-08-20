package handlers

import (
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
	"github.com/taeho-io/family/services/todo_groups/repos/todo_group_permits_repo"
	"github.com/taeho-io/family/services/todo_groups/repos/todo_groups_repo"
)

type ListTodoGroupsFunc func(
	ctx context.Context,
	req *todo_groups.ListTodoGroupsRequest,
) (
	*todo_groups.ListTodoGroupsResponse,
	error,
)

func ListTodoGroups(
	todoGroupsTable *todo_groups_repo.Table,
	todoGroupPermitsTable *todo_group_permits_repo.Table,
) ListTodoGroupsFunc {
	return func(
		ctx context.Context,
		req *todo_groups.ListTodoGroupsRequest,
	) (
		*todo_groups.ListTodoGroupsResponse,
		error,
	) {
		req.AccountId = getAccountIDFromContext(ctx)
		if err := hasPermission(ctx, req.AccountId); err != nil {
			return nil, err
		}

		todoGroupPermits, err := todoGroupPermitsTable.ListByAccountID(req.AccountId)
		if err != nil {
			return nil, err
		}

		var todoGroupIDs []string
		for _, todoGroupPermit := range todoGroupPermits {
			todoGroupIDs = append(todoGroupIDs, todoGroupPermit.TodoGroupID)
		}

		todoGroups, err := todoGroupsTable.ListByIDs(todoGroupIDs)
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
