package repo

import (
	"fmt"
	"time"

	"github.com/guregu/dynamo"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/todos/internal/model"
)

var (
	todoIDFieldKey      = "todo_id"
	parentTypeFieldKey  = "parent_type"
	parentIDFieldKey    = "parent_id"
	parentIDIndexName   = "parent_id-index"
	titleFieldKey       = "title"
	descriptionFieldKey = "description"
	statusFieldKey      = "status"
	assignedToFieldKey  = "assigned_to"
	updatedAtFieldKey   = "updated_at"
	doneAtFieldKey      = "done_at"
	orderFieldKey       = "order"
)

var (
	InvalidTodoError     = fmt.Errorf("invalid todo")
	InvalidTodoIDError   = fmt.Errorf("invliad todo_id")
	InvalidParentIDError = fmt.Errorf("invalid parent_id")
	InvalidTitleError    = fmt.Errorf("invalid title")
)

type TodosRepo interface {
	Put(todo *model.Todo) error
	GetByID(todoID string) (*model.Todo, error)
	ListByIDs(todoIDs []string) ([]*model.Todo, error)
	ListByParentID(parentID string) ([]*model.Todo, error)
	UpdateParent(todoID string, parentType todos.ParentType, parentID string) (*model.Todo, error)
	UpdateTitle(todoID string, title string) (*model.Todo, error)
	UpdateDescription(todoID string, description string) (*model.Todo, error)
	UpdateStatus(todoID string, status todos.Status) (*model.Todo, error)
	UpdateAssignedTo(todoID, accountID string) (*model.Todo, error)
	UpdateOrder(todoID string, order string) (*model.Todo, error)
	DeleteByID(todoID string) error
}

type dynamodbTodosRepo struct {
	TodosRepo
	base.DynamodbRepo

	todosTable *dynamo.Table
}

func NewTodosRepo(ddb base.Dynamodb, cfg TodosRepoConfig) TodosRepo {
	todosTable := ddb.DB().Table(cfg.FullTableName())

	return &dynamodbTodosRepo{
		todosTable: &todosTable,
	}
}

func NewMockTodosRepo() TodosRepo {
	ddb := base.NewMockDynamodb()
	cfg := NewMockTodosRepoConfig()

	return NewTodosRepo(ddb, cfg)
}

func validateTodoInput(todo *model.Todo) error {
	if todo == nil {
		return InvalidTodoError
	}
	if todo.TodoID == "" {
		return InvalidTodoIDError
	}
	if todo.ParentID == "" {
		return InvalidParentIDError
	}
	if todo.Title == "" {
		return InvalidTitleError
	}

	return nil
}

func (t *dynamodbTodosRepo) Table() *dynamo.Table {
	return t.todosTable
}

func (t *dynamodbTodosRepo) Put(todo *model.Todo) error {
	if err := validateTodoInput(todo); err != nil {
		return err
	}

	now := time.Now()
	todo.CreatedAt = now
	todo.UpdatedAt = now

	return t.Table().Put(todo).Run()
}

func (t *dynamodbTodosRepo) GetByID(todoID string) (*model.Todo, error) {
	var todo model.Todo

	if err := t.Table().Get(todoIDFieldKey, todoID).One(&todo); err != nil {
		return nil, err
	}

	return &todo, nil
}

func (t *dynamodbTodosRepo) ListByIDs(todoIDs []string) ([]*model.Todo, error) {
	var todoList []*model.Todo

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

func (t *dynamodbTodosRepo) ListByParentID(parentID string) ([]*model.Todo, error) {
	var todoList []*model.Todo

	err := t.Table().
		Get(parentIDFieldKey, parentID).
		Index(parentIDIndexName).
		All(&todoList)
	if err != nil {
		return nil, err
	}

	return todoList, nil
}

func (t *dynamodbTodosRepo) UpdateParent(todoID string, parentType todos.ParentType, parentID string) (*model.Todo, error) {
	var todo model.Todo

	err := t.Table().
		Update(todoIDFieldKey, todoID).
		If(fmt.Sprintf("%s = ?", todoIDFieldKey), todoID).
		Set(parentTypeFieldKey, parentType).
		Set(parentIDFieldKey, parentID).
		Set(updatedAtFieldKey, time.Now()).
		Value(&todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (t *dynamodbTodosRepo) UpdateTitle(todoID string, title string) (*model.Todo, error) {
	var todo model.Todo

	err := t.Table().
		Update(todoIDFieldKey, todoID).
		If(fmt.Sprintf("%s = ?", todoIDFieldKey), todoID).
		Set(titleFieldKey, title).
		Set(updatedAtFieldKey, time.Now()).
		Value(&todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (t *dynamodbTodosRepo) UpdateDescription(todoID string, description string) (*model.Todo, error) {
	var todo model.Todo

	err := t.Table().
		Update(todoIDFieldKey, todoID).
		If(fmt.Sprintf("%s = ?", todoIDFieldKey), todoID).
		Set(descriptionFieldKey, description).
		Set(updatedAtFieldKey, time.Now()).
		Value(&todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (t *dynamodbTodosRepo) UpdateStatus(todoID string, status todos.Status) (*model.Todo, error) {
	var todo model.Todo

	now := time.Now()
	updateQuery := t.Table().
		Update(todoIDFieldKey, todoID).
		If(fmt.Sprintf("%s = ?", todoIDFieldKey), todoID).
		Set(statusFieldKey, status).
		Set(updatedAtFieldKey, now)

	if status == todos.Status_STATUS_DONE {
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

func (t *dynamodbTodosRepo) UpdateAssignedTo(todoID, accountID string) (*model.Todo, error) {
	var todo model.Todo

	now := time.Now()
	err := t.Table().
		Update(todoIDFieldKey, todoID).
		If(fmt.Sprintf("%s = ?", todoIDFieldKey), todoID).
		Set(assignedToFieldKey, accountID).
		Set(updatedAtFieldKey, now).
		Value(&todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (t *dynamodbTodosRepo) UpdateOrder(todoID string, order string) (*model.Todo, error) {
	var todo model.Todo

	err := t.Table().
		Update(todoIDFieldKey, todoID).
		If(fmt.Sprintf("%s = ?", todoIDFieldKey), todoID).
		Set(orderFieldKey, order).
		Set(updatedAtFieldKey, time.Now()).
		Value(&todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (t *dynamodbTodosRepo) DeleteByID(todoID string) error {
	return t.Table().
		Delete(todoIDFieldKey, todoID).
		If(fmt.Sprintf("%s = ?", todoIDFieldKey), todoID).
		Run()
}
