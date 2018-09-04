package handler

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
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
) UpdateTodoGroupFunc {
	return func(
		ctx context.Context,
		req *todogroups.UpdateTodoGroupRequest,
	) (*todogroups.UpdateTodoGroupResponse, error) {
		accountID := base.GetAccountIDFromContext(ctx)
		if err := base.HasPermissionByAccountID(ctx, accountID); err != nil {
			return nil, err
		}

		logger := ctxlogrus.Extract(ctx).WithField("req", req)

		_, err := todoGroupsRepo.UpdateTitle(
			req.TodoGroup.TodoGroupId,
			req.TodoGroup.Title,
		)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"what": "todoGroupsRepo.UpdateTitle",
			}).Error(err)

			return nil, err
		}

		todoGroup, err := todoGroupsRepo.UpdateDescription(
			req.TodoGroup.TodoGroupId,
			req.TodoGroup.Description,
		)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"what": "todoGroupsRepo.UpdateDescription",
			}).Error(err)

			return nil, err
		}

		return &todogroups.UpdateTodoGroupResponse{
			TodoGroup: todoGroup.ToProto(),
		}, nil
	}
}
