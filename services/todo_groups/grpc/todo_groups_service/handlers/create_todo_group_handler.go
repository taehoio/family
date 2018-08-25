package handlers

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
	"github.com/taeho-io/family/services/base/grpc/base_service"
	"github.com/taeho-io/family/services/todo_groups/models"
	"github.com/taeho-io/family/services/todo_groups/repos/todo_group_permits_repo"
	"github.com/taeho-io/family/services/todo_groups/repos/todo_groups_repo"
)

type CreateTodoGroupFunc func(
	ctx context.Context,
	req *todo_groups.CreateTodoGroupRequest,
) (
	*todo_groups.CreateTodoGroupResponse,
	error,
)

func CreateTodoGroup(
	todoGroupsTable *todo_groups_repo.Table,
	todoGroupPermitsTable *todo_group_permits_repo.Table,
	hasPermissionByAccountID base_service.HasPermissionByAccountIDFunc,
) CreateTodoGroupFunc {
	return func(
		ctx context.Context,
		req *todo_groups.CreateTodoGroupRequest,
	) (
		*todo_groups.CreateTodoGroupResponse,
		error,
	) {
		if err := validateTodoGroupInput(req); err != nil {
			return nil, err
		}

		if err := hasPermissionByAccountID(ctx, req.AccountId); err != nil {
			return nil, err
		}

		logger := ctxlogrus.Extract(ctx)

		todoGroup := models.NewTodoGroupFromProto(req.TodoGroup)
		todoGroup.TodoGroupID = xid.New().String()
		todoGroup.CreatedBy = req.AccountId
		if err := todoGroupsTable.Put(todoGroup); err != nil {
			logger.WithFields(logrus.Fields{
				"what":      "todoGroupsTable.Put",
				"todoGroup": todoGroup,
			}).Error(err)

			return nil, err
		}

		todoGroupPermit := &models.TodoGroupPermit{
			AccountID:   req.AccountId,
			TodoGroupID: todoGroup.TodoGroupID,
			PermitType:  todo_groups.PermitType_PERMIT_TYPE_OWNER,
		}
		if err := todoGroupPermitsTable.Put(todoGroupPermit); err != nil {
			logger.WithFields(logrus.Fields{
				"what":      "todoGroupPermitsTable.Put",
				"todoGroup": todoGroupPermit,
			}).Error(err)

			return nil, err
		}

		return &todo_groups.CreateTodoGroupResponse{
			TodoGroup: todoGroup.ToProto(),
		}, nil
	}
}

func validateTodoGroupInput(req *todo_groups.CreateTodoGroupRequest) error {
	if req.AccountId == "" {
		return base_service.InvalidAccountIDError
	}
	if req.TodoGroup.Title == "" {
		return InvalidTitleError
	}

	return nil
}
