package todos_repo

import (
	"fmt"
	"strings"
	"time"

	"github.com/guregu/dynamo"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
	"github.com/taeho-io/family/services/base/aws/dynamodb"
	"github.com/taeho-io/family/services/base/aws/dynamodb/table"
	"github.com/taeho-io/family/services/todos/config"
	"github.com/taeho-io/family/services/todos/models"
)

const (
	todoIDFieldKey      = "todo_id"
	titleFieldKey       = "title"
	statusFieldKey      = "status"
	descriptionFieldKey = "description"
	updatedAtFieldKey   = "updated_at"
	doneAtFieldKey      = "done_at"
)

var (
	InvalidTodoError        = fmt.Errorf("invalid todo")
	InvalidTodoIDError      = fmt.Errorf("invliad todo_id")
	InvalidTodoGroupIDError = fmt.Errorf("invalid todo_group_id")
)

type Table struct {
	table.IFace

	todosTable *dynamo.Table
}

func New(ddb dynamodb.IFace, cfg config.IFace) *Table {
	fullTableName := fullTableName(cfg)
	todosTable := ddb.DB().Table(fullTableName)

	return &Table{
		todosTable: &todosTable,
	}
}

func NewMock() *Table {
	ddb := dynamodb.NewMock()
	cfg := config.NewMock()

	return New(ddb, cfg)
}

func fullTableName(cfg config.IFace) string {
	prefix := cfg.Prefix()
	tableName := cfg.Settings().DynamodbTodosTableName
	return strings.Join([]string{prefix, tableName}, "-")
}

func (t *Table) Table() *dynamo.Table {
	return t.todosTable
}

func (t *Table) validateTodoInput(todo *models.Todo) error {
	if todo == nil {
		return InvalidTodoError
	}
	if todo.TodoID == "" {
		return InvalidTodoIDError
	}
	if todo.TodoGroupID == "" {
		return InvalidTodoGroupIDError
	}

	return nil
}

func (t *Table) Put(todo *models.Todo) error {
	if err := t.validateTodoInput(todo); err != nil {
		return err
	}

	now := time.Now()
	todo.CreatedAt = now
	todo.UpdatedAt = now

	return t.Table().Put(todo).Run()
}

func (t *Table) GetByID(todoID string) (*models.Todo, error) {
	var todo models.Todo

	if err := t.Table().Get(todoIDFieldKey, todoID).One(&todo); err != nil {
		return nil, err
	}

	return &todo, nil
}

func (t *Table) ListByIDs(todoIDs []string) ([]*models.Todo, error) {
	var todoList []*models.Todo

	// TODO: make it concurrent.
	for _, todoID := range todoIDs {
		todo, err := t.GetByID(todoID)
		if err != nil {
			return nil, err
		}
		todoList = append(todoList, todo)
	}

	return todoList, nil
}

func (t *Table) UpdateTitle(todoID string, title string) (*models.Todo, error) {
	var todo models.Todo

	err := t.Table().
		Update(todoIDFieldKey, todoID).
		If("todo_id = ?", todoID).
		Set(titleFieldKey, title).
		Set(updatedAtFieldKey, time.Now()).
		Value(&todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (t *Table) UpdateDescription(todoID string, description string) (*models.Todo, error) {
	var todo models.Todo

	err := t.Table().
		Update(todoIDFieldKey, todoID).
		If("todo_id = ?", todoID).
		Set(descriptionFieldKey, description).
		Set(updatedAtFieldKey, time.Now()).
		Value(&todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (t *Table) UpdateStatus(todoID string, status string) (*models.Todo, error) {
	var todo models.Todo

	now := time.Now()
	updateQuery := t.Table().
		Update(todoIDFieldKey, todoID).
		If("todo_id = ?", todoID).
		Set(statusFieldKey, status).
		Set(updatedAtFieldKey, now)

	if status == todos.Status_DONE.String() {
		updateQuery.Set(doneAtFieldKey, now)
	} else {
		updateQuery.Set(doneAtFieldKey, time.Unix(0, 0))
	}

	err := updateQuery.Value(&todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (t *Table) DeleteByID(todoID string) error {
	return t.Table().
		Delete(todoIDFieldKey, todoID).
		If(fmt.Sprintf("%s = ?", todoIDFieldKey), todoID).
		Run()
}
