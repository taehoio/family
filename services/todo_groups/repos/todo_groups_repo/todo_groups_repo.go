package todo_groups_repo

import (
	"fmt"
	"strings"
	"time"

	"github.com/taeho-io/family/services/base/aws/dynamodb/table"

	"github.com/guregu/dynamo"

	"github.com/taeho-io/family/services/base/aws/dynamodb"
	"github.com/taeho-io/family/services/todo_groups/config"
	"github.com/taeho-io/family/services/todo_groups/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

type IFace interface {
	table.IFace

	Put(*models.TodoGroup) error
	GetByID(string) (*models.TodoGroup, error)
	ListByIDs([]string) ([]*models.TodoGroup, error)
	UpdateTitle(string, string) (*models.TodoGroup, error)
	UpdateDescription(string, string) (*models.TodoGroup, error)
	UpdateOrder(string, string) (*models.TodoGroup, error)
	DeleteByID(string) error
}

type Table struct {
	IFace

	todoGroupsTable *dynamo.Table
}

func New(ddb dynamodb.IFace, cfg config.IFace) *Table {
	fullTableName := fullTableName(cfg)
	todoGroupsTable := ddb.DB().Table(fullTableName)

	return &Table{
		todoGroupsTable: &todoGroupsTable,
	}
}

func NewMock() *Table {
	ddb := dynamodb.NewMock()
	cfg := config.NewMock()

	return New(ddb, cfg)
}

func fullTableName(cfg config.IFace) string {
	prefix := cfg.Prefix()
	tableName := cfg.Settings().DynamodbTodoGroupsTableName
	return strings.Join([]string{prefix, tableName}, "-")
}

func validateTodoGroupInput(todoGroup *models.TodoGroup) error {
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

func (t *Table) Table() *dynamo.Table {
	return t.todoGroupsTable
}

func (t *Table) Put(todoGroup *models.TodoGroup) error {
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

func (t *Table) GetByID(todoGroupID string) (*models.TodoGroup, error) {
	var todoGroup models.TodoGroup

	if err := t.Table().Get(todoGroupIDFieldKey, todoGroupID).One(&todoGroup); err != nil {
		return nil, err
	}

	return &todoGroup, nil
}

func (t *Table) ListByIDs(todoGroupIDs []string) ([]*models.TodoGroup, error) {
	var todoGroups []*models.TodoGroup

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

func (t *Table) UpdateTitle(todoGroupID, title string) (*models.TodoGroup, error) {
	var todoGroup models.TodoGroup

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

func (t *Table) UpdateDescription(todoGroupID, description string) (*models.TodoGroup, error) {
	var todoGroup models.TodoGroup

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

func (t *Table) UpdateOrder(todoGroupID, order string) (*models.TodoGroup, error) {
	var todoGroup models.TodoGroup

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

func (t *Table) DeleteByID(todoGroupID string) error {
	return t.Table().
		Delete(todoGroupIDFieldKey, todoGroupID).
		If(fmt.Sprintf("%s = ?", todoGroupIDFieldKey), todoGroupID).
		Run()
}
