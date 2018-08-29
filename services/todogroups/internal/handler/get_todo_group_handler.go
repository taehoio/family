package handler

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/guregu/dynamo"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todogroups"
	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/todogroups/internal/repo"
)

type GetTodoGroupFunc func(
	ctx context.Context,
	req *todogroups.GetTodoGroupRequest,
) (
	*todogroups.GetTodoGroupResponse,
	error,
)

func GetTodoGroup(
	todoGroupsRepo repo.GroupsRepo,
	todoGroupPermitsRepo repo.PermitsRepo,
	getAccountIDFromContext base.GetAccountIDFromContextFunc,
	hasPermissionByAccountID base.HasPermissionByAccountIDFunc,
) GetTodoGroupFunc {
	return func(
		ctx context.Context,
		req *todogroups.GetTodoGroupRequest,
	) (
		*todogroups.GetTodoGroupResponse,
		error,
	) {
		req.AccountId = getAccountIDFromContext(ctx)
		if err := hasPermissionByAccountID(ctx, req.AccountId); err != nil {
			return nil, err
		}

		if req.AccountId == "" {
			return nil, base.InvalidAccountIDError
		}
		if req.TodoGroupId == "" {
			return nil, InvalidTodoGroupIDError
		}

		logger := ctxlogrus.Extract(ctx).WithField("req", req)

		togoGroupPermit, err := todoGroupPermitsRepo.Get(req.AccountId, req.TodoGroupId)
		if err == dynamo.ErrNotFound {
			logger.WithFields(logrus.Fields{
				"what": "todoGroupPermitsRepo.Get",
				"req":  req,
			}).Warn(err)

			return nil, base.PermissionDeniedError
		}
		if err != nil {
			logger.WithFields(logrus.Fields{
				"what": "todoGroupPermitsRepo.Get",
				"req":  req,
			}).Error(err)

			return nil, base.PermissionDeniedError
		}

		todoGroup, err := todoGroupsRepo.GetByID(req.TodoGroupId)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"what": "todoGroupsRepo.GetByID",
				"req":  req,
			}).Error(err)

			return nil, err
		}
		todoGroup.PermitType = togoGroupPermit.PermitType

		return &todogroups.GetTodoGroupResponse{
			TodoGroup: todoGroup.ToProto(),
		}, nil
	}
}
