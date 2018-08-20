package handlers

import (
	"github.com/guregu/dynamo"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
	"github.com/taeho-io/family/services/todo_groups/repos/todo_group_permits_repo"
	"github.com/taeho-io/family/services/todo_groups/repos/todo_groups_repo"
)

type GetTodoGroupFunc func(
	ctx context.Context,
	req *todo_groups.GetTodoGroupRequest,
) (
	*todo_groups.GetTodoGroupResponse,
	error,
)

func GetTodoGroup(
	logger *logrus.Entry,
	todoGroupsTable *todo_groups_repo.Table,
	todoGroupPermitsTable *todo_group_permits_repo.Table,
) GetTodoGroupFunc {
	return func(
		ctx context.Context,
		req *todo_groups.GetTodoGroupRequest,
	) (
		*todo_groups.GetTodoGroupResponse,
		error,
	) {
		req.AccountId = getAccountIDFromContext(ctx)
		if req.AccountId == "" {
			return nil, InvalidAccountIDError
		}
		if req.TodoGroupId == "" {
			return nil, InvalidTodoGroupIDError
		}

		if err := hasPermission(ctx, req.AccountId); err != nil {
			return nil, err
		}

		_, err := todoGroupPermitsTable.Get(req.AccountId, req.TodoGroupId)
		if err == dynamo.ErrNotFound {
			logger.Warn(err)
			return nil, PermissionDeniedError
		}
		if err != nil {
			logger.Error(err)
			return nil, PermissionDeniedError
		}

		todoGroup, err := todoGroupsTable.GetByID(req.TodoGroupId)
		if err != nil {
			return nil, err
		}

		return &todo_groups.GetTodoGroupResponse{
			TodoGroup: todoGroup.ToProto(),
		}, nil
	}
}
