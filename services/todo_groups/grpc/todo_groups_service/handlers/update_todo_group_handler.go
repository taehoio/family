package handlers

import (
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
	"github.com/taeho-io/family/services/base/grpc/base_service"
	"github.com/taeho-io/family/services/todo_groups/repos/todo_groups_repo"
)

type UpdateTodoGroupFunc func(
	ctx context.Context, req *todo_groups.UpdateTodoGroupRequest,
) (*todo_groups.UpdateTodoGroupResponse, error)

func UpdateTodoGroup(
	todoGroupsTable todo_groups_repo.IFace,
	hasPermissionByAccountID base_service.HasPermissionByAccountIDFunc,
) UpdateTodoGroupFunc {
	return func(
		ctx context.Context,
		req *todo_groups.UpdateTodoGroupRequest,
	) (*todo_groups.UpdateTodoGroupResponse, error) {
		if err := hasPermissionByAccountID(ctx, req.AccountId); err != nil {
			return nil, err
		}

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
