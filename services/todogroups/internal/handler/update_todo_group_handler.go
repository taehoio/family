package handler

import (
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todogroups"
	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/todogroups/internal/repo"
)

type UpdateTodoGroupFunc func(
	ctx context.Context, req *todogroups.UpdateTodoGroupRequest,
) (*todogroups.UpdateTodoGroupResponse, error)

func UpdateTodoGroup(
	todoGroupsRepo repo.GroupsRepo,
	hasPermissionByAccountID base.HasPermissionByAccountIDFunc,
) UpdateTodoGroupFunc {
	return func(
		ctx context.Context,
		req *todogroups.UpdateTodoGroupRequest,
	) (*todogroups.UpdateTodoGroupResponse, error) {
		if err := hasPermissionByAccountID(ctx, req.AccountId); err != nil {
			return nil, err
		}

		_, err := todoGroupsRepo.UpdateTitle(
			req.GetTodoGroupId(),
			req.GetTodoGroup().GetTitle(),
		)
		if err != nil {
			return nil, err
		}

		todoGroup, err := todoGroupsRepo.UpdateDescription(
			req.GetTodoGroupId(),
			req.GetTodoGroup().GetDescription(),
		)
		if err != nil {
			return nil, err
		}

		return &todogroups.UpdateTodoGroupResponse{
			TodoGroup: todoGroup.ToProto(),
		}, nil
	}
}
