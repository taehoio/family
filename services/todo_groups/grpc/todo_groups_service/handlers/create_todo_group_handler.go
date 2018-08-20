package handlers

import (
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
	"github.com/taeho-io/family/services/base/grpc/interceptors"
	"github.com/taeho-io/family/services/todo_groups/models"
	"github.com/taeho-io/family/services/todo_groups/repos/todo_group_permits_repo"
	"github.com/taeho-io/family/services/todo_groups/repos/todo_groups_repo"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateTodoGroupFunc func(
	ctx context.Context,
	req *todo_groups.CreateTodoGroupRequest,
) (
	*todo_groups.CreateTodoGroupResponse,
	error,
)

func CreateTodoGroup(
	logger *logrus.Entry,
	todoGroupsTable *todo_groups_repo.Table,
	todoGroupPermitsTable *todo_group_permits_repo.Table,
) CreateTodoGroupFunc {
	return func(
		ctx context.Context,
		req *todo_groups.CreateTodoGroupRequest,
	) (
		*todo_groups.CreateTodoGroupResponse,
		error,
	) {
		err := validateTodoGroupInput(req)
		if err != nil {
			return nil, err
		}

		if err := hasPermission(ctx, req.AccountId); err != nil {
			return nil, err
		}

		todoGroup := models.NewTodoGroupFromProto(req.GetTodoGroup())
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
			PermitType:  todo_groups.TodoGroupPermitType_OWNER.String(),
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
		return InvalidAccountIDError
	}
	if req.TodoGroup.Title == "" {
		return InvalidTitleError
	}

	return nil
}

func getAccountIDFromContext(ctx context.Context) string {
	return ctx.Value(interceptors.AccountIDKey).(string)
}

func hasPermission(ctx context.Context, accountID string) error {
	if accountID == "" {
		return InvalidAccountIDError
	}

	accountIDFromCtx := ctx.Value(interceptors.AccountIDKey)
	if accountIDFromCtx != accountID {
		return status.Error(codes.PermissionDenied, "Forbidden")
	}

	return nil
}
