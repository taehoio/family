package handler

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todogroups"
	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/todogroups/internal/repo"
)

type ListTodoGroupsFunc func(
	ctx context.Context,
	req *todogroups.ListTodoGroupsRequest,
) (
	*todogroups.ListTodoGroupsResponse,
	error,
)

func ListTodoGroups(
	todoGroupsRepo repo.TodoGroupsRepo,
	todoGroupPermitsRepo repo.TodoGroupPermitsRepo,
	hasPermissionByAccountID base.HasPermissionByAccountIDFunc,
) ListTodoGroupsFunc {
	return func(
		ctx context.Context,
		req *todogroups.ListTodoGroupsRequest,
	) (
		*todogroups.ListTodoGroupsResponse,
		error,
	) {
		if err := hasPermissionByAccountID(ctx, req.AccountId); err != nil {
			return nil, err
		}

		logger := ctxlogrus.Extract(ctx)

		todoGroupPermits, err := todoGroupPermitsRepo.ListByAccountID(req.AccountId)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"what":      "todoGroupPermitsRepo.ListByAccountID",
				"accountId": req.AccountId,
			}).Error(err)

			return nil, err
		}

		var todoGroupIDs []string
		for _, todoGroupPermit := range todoGroupPermits {
			todoGroupIDs = append(todoGroupIDs, todoGroupPermit.TodoGroupID)
		}

		todoGroups, err := todoGroupsRepo.ListByIDs(todoGroupIDs)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"what":         "todoGroupsRepo.ListByIDs",
				"todoGroupIDs": todoGroupIDs,
			}).Error(err)

			return nil, err
		}

		var todoGroupProtos []*todogroups.TodoGroup
		for _, todoGroup := range todoGroups {
			todoGroupProtos = append(todoGroupProtos, todoGroup.ToProto())
		}

		return &todogroups.ListTodoGroupsResponse{
			TodoGroups: todoGroupProtos,
		}, nil
	}
}