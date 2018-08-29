package handler

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todogroups"
	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/todogroups/internal/model"
	"github.com/taeho-io/family/services/todogroups/internal/repo"
)

type CreateTodoGroupFunc func(
	ctx context.Context,
	req *todogroups.CreateTodoGroupRequest,
) (
	*todogroups.CreateTodoGroupResponse,
	error,
)

func CreateTodoGroup(
	todoGroupsRepo repo.TodoGroupsRepo,
	todoGroupPermitsRepo repo.TodoGroupPermitsRepo,
	hasPermissionByAccountID base.HasPermissionByAccountIDFunc,
) CreateTodoGroupFunc {
	return func(
		ctx context.Context,
		req *todogroups.CreateTodoGroupRequest,
	) (
		*todogroups.CreateTodoGroupResponse,
		error,
	) {
		if err := validateTodoGroupInput(req); err != nil {
			return nil, err
		}

		if err := hasPermissionByAccountID(ctx, req.AccountId); err != nil {
			return nil, err
		}

		logger := ctxlogrus.Extract(ctx)

		todoGroup := model.NewTodoGroupFromProto(req.TodoGroup)
		todoGroup.TodoGroupID = xid.New().String()
		todoGroup.CreatedBy = req.AccountId
		if err := todoGroupsRepo.Put(todoGroup); err != nil {
			logger.WithFields(logrus.Fields{
				"what":      "todoGroupsRepo.Put",
				"todoGroup": todoGroup,
			}).Error(err)

			return nil, err
		}

		todoGroupPermit := &model.TodoGroupPermit{
			AccountID:   req.AccountId,
			TodoGroupID: todoGroup.TodoGroupID,
			PermitType:  todogroups.PermitType_PERMIT_TYPE_OWNER,
		}
		if err := todoGroupPermitsRepo.Put(todoGroupPermit); err != nil {
			logger.WithFields(logrus.Fields{
				"what":      "todoGroupPermitsRepo.Put",
				"todoGroup": todoGroupPermit,
			}).Error(err)

			return nil, err
		}

		return &todogroups.CreateTodoGroupResponse{
			TodoGroup: todoGroup.ToProto(),
		}, nil
	}
}

func validateTodoGroupInput(req *todogroups.CreateTodoGroupRequest) error {
	if req.AccountId == "" {
		return base.InvalidAccountIDError
	}
	if req.TodoGroup.Title == "" {
		return InvalidTitleError
	}

	return nil
}
