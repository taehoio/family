package handlers

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
	"github.com/taeho-io/family/services/base/grpc/base_service"
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
	hasPermissionByAccountID base_service.HasPermissionByAccountIDFunc,
) ListTodoGroupsFunc {
	return func(
		ctx context.Context,
		req *todo_groups.ListTodoGroupsRequest,
	) (
		*todo_groups.ListTodoGroupsResponse,
		error,
	) {
		if err := hasPermissionByAccountID(ctx, req.AccountId); err != nil {
			return nil, err
		}

		logger := ctxlogrus.Extract(ctx)

		todoGroupPermits, err := todoGroupPermitsTable.ListByAccountID(req.AccountId)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"what":      "todoGroupPermitsTable.ListByAccountID",
				"accountId": req.AccountId,
			}).Error(err)

			return nil, err
		}

		var todoGroupIDs []string
		for _, todoGroupPermit := range todoGroupPermits {
			todoGroupIDs = append(todoGroupIDs, todoGroupPermit.TodoGroupID)
		}

		todoGroups, err := todoGroupsTable.ListByIDs(todoGroupIDs)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"what":         "todoGroupsTable.ListByIDs",
				"todoGroupIDs": todoGroupIDs,
			}).Error(err)

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
