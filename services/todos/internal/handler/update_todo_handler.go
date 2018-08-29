package handler

import (
	"fmt"

	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/todos/internal/repo"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/taeho-io/family/idl/generated/go/pb/family/todogroups"
	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
	"github.com/taeho-io/family/services/todos/internal/model"
)

type UpdateTodoFunc func(ctx context.Context, req *todos.UpdateTodoRequest) (*todos.UpdateTodoResponse, error)

func UpdateTodo(
	todosRepo repo.TodosRepo,
	getAccountIDFromContext base.GetAccountIDFromContextFunc,
	hasPermissionByAccountID base.HasPermissionByAccountIDFunc,
	todoGroupsServiceClient todogroups.TodoGroupsServiceClient,
) UpdateTodoFunc {
	return func(ctx context.Context, req *todos.UpdateTodoRequest) (*todos.UpdateTodoResponse, error) {
		req.AccountId = getAccountIDFromContext(ctx)
		if err := hasPermissionByAccountID(ctx, req.AccountId); err != nil {
			return nil, err
		}

		logger := ctxlogrus.Extract(ctx).WithField("req", req)

		getTogoGroupReq := &todogroups.GetTodoGroupRequest{
			AccountId:   req.AccountId,
			TodoGroupId: req.Todo.ParentId,
		}
		todoGroupRes, err := todoGroupsServiceClient.GetTodoGroup(ctx, getTogoGroupReq)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"what":            "todoGroupsServiceClient.GetTodoGroup",
				"getTogoGroupReq": getTogoGroupReq,
			}).Error(err)

			return nil, err
		}
		if todoGroupRes.TodoGroup == nil {
			return nil, base.PermissionDeniedError
		}

		todo := &model.Todo{}

		for _, field := range req.Fields {
			switch field {
			case todos.UpdatingField_UPDATING_FIELD_PARENT:
				todo, err = todosRepo.UpdateParent(req.TodoId, req.Todo.ParentType, req.Todo.ParentId)
				if err != nil {
					return nil, err
				}
			case todos.UpdatingField_UPDATING_FIELD_TITLE:
				todo, err = todosRepo.UpdateTitle(req.TodoId, req.Todo.Title)
				if err != nil {
					return nil, err
				}
			case todos.UpdatingField_UPDATING_FIELD_DESCRIPTION:
				todo, err = todosRepo.UpdateDescription(req.TodoId, req.Todo.Description)
				if err != nil {
					return nil, err
				}
			case todos.UpdatingField_UPDATING_FIELD_STATUS:
				todo, err = todosRepo.UpdateStatus(req.TodoId, req.Todo.Status)
				if err != nil {
					return nil, err
				}
			case todos.UpdatingField_UPDATING_FIELD_ORDER:
				todo, err = todosRepo.UpdateOrder(req.TodoId, req.Todo.Order)
				if err != nil {
					return nil, err
				}
			case todos.UpdatingField_UPDATING_FIELD_ASSIGNED_TO:
				todo, err = todosRepo.UpdateAssignedTo(req.TodoId, req.Todo.AssignedTo)
				if err != nil {
					return nil, err
				}
			default:
				logger.Warn(fmt.Errorf("invalid updating field %s", field.String()))
			}
		}

		return &todos.UpdateTodoResponse{
			Todo: todo.ToProto(),
		}, nil
	}
}
