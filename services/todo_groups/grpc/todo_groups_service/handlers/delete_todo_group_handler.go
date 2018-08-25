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

type DeleteTodoGroupFunc func(
	ctx context.Context,
	req *todo_groups.DeleteTodoGroupRequest,
) (*todo_groups.DeleteTodoGroupResponse, error)

func DeleteTodoGroup(
	todoGroupsTable *todo_groups_repo.Table,
	todoGroupPermitsTable *todo_group_permits_repo.Table,
	hasPermissionByAccountID base_service.HasPermissionByAccountIDFunc,
) DeleteTodoGroupFunc {
	return func(
		ctx context.Context,
		req *todo_groups.DeleteTodoGroupRequest,
	) (*todo_groups.DeleteTodoGroupResponse, error) {
		if err := hasPermissionByAccountID(ctx, req.AccountId); err != nil {
			return nil, err
		}

		logger := ctxlogrus.Extract(ctx)

		if err := todoGroupsTable.DeleteByID(req.TodoGroupId); err != nil {
			logger.WithFields(logrus.Fields{
				"what": "todoGroupsTable.DeleteByID",
				"req":  req,
			}).Error(err)

			return nil, err
		}

		if err := todoGroupPermitsTable.Delete(req.AccountId, req.TodoGroupId); err != nil {
			logger.WithFields(logrus.Fields{
				"what": "todoGroupPermitsTable.Delete",
				"req":  req,
			}).Error(err)

			return nil, err
		}

		return &todo_groups.DeleteTodoGroupResponse{}, nil
	}
}
