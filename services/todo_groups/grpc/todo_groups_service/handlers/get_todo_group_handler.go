package handlers

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/guregu/dynamo"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
	"github.com/taeho-io/family/services/base/grpc/base_service"
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
	todoGroupsTable *todo_groups_repo.Table,
	todoGroupPermitsTable *todo_group_permits_repo.Table,
	getAccountIDFromContext base_service.GetAccountIDFromContextFunc,
	hasPermissionByAccountID base_service.HasPermissionByAccountIDFunc,
) GetTodoGroupFunc {
	return func(
		ctx context.Context,
		req *todo_groups.GetTodoGroupRequest,
	) (
		*todo_groups.GetTodoGroupResponse,
		error,
	) {
		req.AccountId = getAccountIDFromContext(ctx)
		if err := hasPermissionByAccountID(ctx, req.AccountId); err != nil {
			return nil, err
		}

		if req.AccountId == "" {
			return nil, base_service.InvalidAccountIDError
		}
		if req.TodoGroupId == "" {
			return nil, InvalidTodoGroupIDError
		}

		logger := ctxlogrus.Extract(ctx).WithField("req", req)

		togoGroupPermit, err := todoGroupPermitsTable.Get(req.AccountId, req.TodoGroupId)
		if err == dynamo.ErrNotFound {
			logger.WithFields(logrus.Fields{
				"what": "todoGroupPermitsTable.Get",
				"req":  req,
			}).Warn(err)

			return nil, base_service.PermissionDeniedError
		}
		if err != nil {
			logger.WithFields(logrus.Fields{
				"what": "todoGroupPermitsTable.Get",
				"req":  req,
			}).Error(err)

			return nil, base_service.PermissionDeniedError
		}

		todoGroup, err := todoGroupsTable.GetByID(req.TodoGroupId)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"what": "todoGroupsTable.GetByID",
				"req":  req,
			}).Error(err)

			return nil, err
		}
		todoGroup.PermitType = togoGroupPermit.PermitType

		return &todo_groups.GetTodoGroupResponse{
			TodoGroup: todoGroup.ToProto(),
		}, nil
	}
}
