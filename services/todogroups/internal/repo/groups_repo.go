package repo

import (
	"fmt"
	"time"

	"github.com/guregu/dynamo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/todogroups/internal/model"
)

const (
	todoGroupIDFieldKey = "todo_group_id"
	titleFieldKey       = "title"
	descriptionFieldKey = "description"
	updatedAtFieldKey   = "updated_at"
	orderFieldKey       = "order"
)

var (
	InvalidTodoGroupIDError = status.Error(codes.InvalidArgument, "invalid todo_group_id")
	InvalidTitleError       = status.Error(codes.InvalidArgument, "invalid title")
	InvalidCreatedByError   = status.Error(codes.InvalidArgument, "invalid created_by")
)

type GroupsRepo interface {
	Put(*model.TodoGroup) error
	GetByID(string) (*model.TodoGroup, error)
	ListByIDs([]string) ([]*model.TodoGroup, error)
	UpdateTitle(string, string) (*model.TodoGroup, error)
	UpdateDescription(string, string) (*model.TodoGroup, error)
	UpdateOrder(string, string) (*model.TodoGroup, error)
	DeleteByID(string) error
}

type dynamodbGroupsRepo struct {
	GroupsRepo
	base.DynamodbRepo

	todoGroupsTable *dynamo.Table
}

func NewTodosRepo(ddb base.Dynamodb, cfg TodoGroupsRepoConfig) GroupsRepo {
	todoGroupsTable := ddb.DB().Table(cfg.FullTableName())

	return &dynamodbGroupsRepo{
		todoGroupsTable: &todoGroupsTable,
	}
}

func NewMockGroupsRepo() GroupsRepo {
	ddb := base.NewMockDynamodb()
	cfg := NewMockTodoGroupsRepoConfig()

	return NewTodosRepo(ddb, cfg)
}

func validateTodoGroupInput(todoGroup *model.TodoGroup) error {
	if todoGroup.TodoGroupID == "" {
		return InvalidTodoGroupIDError
	}
	if todoGroup.Title == "" {
		return InvalidTitleError
	}
	if todoGroup.CreatedBy == "" {
		return InvalidCreatedByError
	}

	return nil
}

func (t *dynamodbGroupsRepo) Table() *dynamo.Table {
	return t.todoGroupsTable
}

func (t *dynamodbGroupsRepo) Put(todoGroup *model.TodoGroup) error {
	if err := validateTodoGroupInput(todoGroup); err != nil {
		return err
	}

	now := time.Now()
	todoGroup.CreatedAt = now
	todoGroup.UpdatedAt = now

	return t.Table().
		Put(todoGroup).
		Run()
}

func (t *dynamodbGroupsRepo) GetByID(todoGroupID string) (*model.TodoGroup, error) {
	var todoGroup model.TodoGroup

	if err := t.Table().Get(todoGroupIDFieldKey, todoGroupID).One(&todoGroup); err != nil {
		return nil, err
	}

	return &todoGroup, nil
}

func (t *dynamodbGroupsRepo) ListByIDs(todoGroupIDs []string) ([]*model.TodoGroup, error) {
	var todoGroups []*model.TodoGroup

	// TODO: make it concurrent.
	for _, todoGroupID := range todoGroupIDs {
		todoGroup, err := t.GetByID(todoGroupID)
		if err != nil {
			return nil, err
		}
		todoGroups = append(todoGroups, todoGroup)
	}

	return todoGroups, nil
}

func (t *dynamodbGroupsRepo) UpdateTitle(todoGroupID, title string) (*model.TodoGroup, error) {
	var todoGroup model.TodoGroup

	err := t.Table().
		Update(todoGroupIDFieldKey, todoGroupID).
		If(fmt.Sprintf("%s = ?", todoGroupIDFieldKey), todoGroupID).
		Set(titleFieldKey, title).
		Set(updatedAtFieldKey, time.Now()).
		Value(&todoGroup)
	if err != nil {
		return nil, err
	}

	return &todoGroup, nil
}

func (t *dynamodbGroupsRepo) UpdateDescription(todoGroupID, description string) (*model.TodoGroup, error) {
	var todoGroup model.TodoGroup

	err := t.Table().
		Update(todoGroupIDFieldKey, todoGroupID).
		If(fmt.Sprintf("%s = ?", todoGroupIDFieldKey), todoGroupID).
		Set(descriptionFieldKey, description).
		Set(updatedAtFieldKey, time.Now()).
		Value(&todoGroup)
	if err != nil {
		return nil, err
	}

	return &todoGroup, nil
}

func (t *dynamodbGroupsRepo) UpdateOrder(todoGroupID, order string) (*model.TodoGroup, error) {
	var todoGroup model.TodoGroup

	err := t.Table().
		Update(todoGroupIDFieldKey, todoGroupID).
		If(fmt.Sprintf("%s = ?", todoGroupIDFieldKey), todoGroupID).
		Set(orderFieldKey, order).
		Set(updatedAtFieldKey, time.Now()).
		Value(&todoGroup)
	if err != nil {
		return nil, err
	}

	return &todoGroup, nil
}

func (t *dynamodbGroupsRepo) DeleteByID(todoGroupID string) error {
	return t.Table().
		Delete(todoGroupIDFieldKey, todoGroupID).
		If(fmt.Sprintf("%s = ?", todoGroupIDFieldKey), todoGroupID).
		Run()
}
