package todos_repo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
	"github.com/taeho-io/family/services/todos/models"
)

var (
	todosTable *Table

	testFullTableName      = "family-development-todos-todos"
	testTodoID             = "test_todo_id"
	testNonExistTodoID     = "test_no_exist_todo_id"
	testTodoGroupID        = "test_todo_group_id"
	testTodoTypeTodo       = todos.Status_DONE.String()
	testTitle              = "test_title"
	testDescription        = "test_description"
	testUpdatedTitle       = "test_updated_title"
	testUpdatedDescription = "test_updated_description"
	testUpdatedOrder       = "m"
)

func TestMain(m *testing.M) {
	todosTable = NewMock()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestFullTableName(t *testing.T) {
	assert.Equal(t, testFullTableName, todosTable.Table().Name())
}

func TestValidateTodoInput(t *testing.T) {
	todo := &models.Todo{
		TodoID:      testTodoID,
		TodoGroupID: testTodoGroupID,
		Title:       testTitle,
		Description: testDescription,
		Status:      testTodoTypeTodo,
	}
	err := todosTable.validateTodoInput(todo)
	assert.Nil(t, err)
}

func TestValidateTodoInputInvalidTodo(t *testing.T) {
	err := todosTable.validateTodoInput(nil)
	assert.Equal(t, InvalidTodoError, err)
}

func TestValidateTodoInputInvalidTodoID(t *testing.T) {
	todo := &models.Todo{
		TodoGroupID: testTodoGroupID,
		Title:       testTitle,
		Description: testDescription,
		Status:      testTodoTypeTodo,
	}
	err := todosTable.validateTodoInput(todo)
	assert.Equal(t, InvalidTodoIDError, err)
}

func TestValidateTodoInputInvalidTodoGroupID(t *testing.T) {
	todo := &models.Todo{
		TodoID:      testTodoID,
		Title:       testTitle,
		Description: testDescription,
		Status:      testTodoTypeTodo,
	}
	err := todosTable.validateTodoInput(todo)
	assert.Equal(t, InvalidTodoGroupIDError, err)
}

func TestValidateTodoInputInvalidTitle(t *testing.T) {
	todo := &models.Todo{
		TodoID:      testTodoID,
		TodoGroupID: testTodoGroupID,
		Description: testDescription,
		Status:      testTodoTypeTodo,
	}
	err := todosTable.validateTodoInput(todo)
	assert.Equal(t, InvalidTitleError, err)
}

func TestPut(t *testing.T) {
	todo := &models.Todo{
		TodoID:      testTodoID,
		TodoGroupID: testTodoGroupID,
		Title:       testTitle,
		Description: testDescription,
		Status:      testTodoTypeTodo,
	}
	err := todosTable.Put(todo)
	assert.Nil(t, err)
}

func TestPutInvalidTodo(t *testing.T) {
	err := todosTable.Put(nil)
	assert.NotNil(t, err)
}

func TestGetByID(t *testing.T) {
	todo, err := todosTable.GetByID(testTodoID)
	assert.Nil(t, err)
	assert.NotNil(t, todo)
	assert.Equal(t, testTodoGroupID, todo.TodoGroupID)
	assert.Equal(t, testTitle, todo.Title)
}

func TestGetByIDFail(t *testing.T) {
	todo, err := todosTable.GetByID(testNonExistTodoID)
	assert.NotNil(t, err)
	assert.Nil(t, todo)
}

func TestListByIDsFail(t *testing.T) {
	todoIDs := []string{testNonExistTodoID}
	todoList, err := todosTable.ListByIDs(todoIDs)
	assert.NotNil(t, err)
	assert.Nil(t, todoList)
}

func TestListByIDs(t *testing.T) {
	todoIDs := []string{testTodoID}
	todoList, err := todosTable.ListByIDs(todoIDs)
	assert.Nil(t, err)
	assert.NotNil(t, todoList)
	assert.Len(t, todoList, 1)

	todo := todoList[0]
	assert.Equal(t, testTodoGroupID, todo.TodoGroupID)
	assert.Equal(t, testTitle, todo.Title)
}

func TestUpdateTitleFail(t *testing.T) {
	todo, err := todosTable.UpdateTitle(testNonExistTodoID, testUpdatedTitle)
	assert.Nil(t, todo)
	assert.NotNil(t, err)
}

func TestUpdateTitle(t *testing.T) {
	todo, err := todosTable.UpdateTitle(testTodoID, testUpdatedTitle)
	assert.NotNil(t, todo)
	assert.Nil(t, err)
	assert.Equal(t, testUpdatedTitle, todo.Title)
}

func TestUpdateDescriptionFail(t *testing.T) {
	todo, err := todosTable.UpdateDescription(testNonExistTodoID, testUpdatedDescription)
	assert.Nil(t, todo)
	assert.NotNil(t, err)
}

func TestUpdateDescription(t *testing.T) {
	todo, err := todosTable.UpdateDescription(testTodoID, testUpdatedDescription)
	assert.NotNil(t, todo)
	assert.Nil(t, err)
	assert.Equal(t, testUpdatedDescription, todo.Description)
}

func TestUpdateTodoTypeFail(t *testing.T) {
	todo, err := todosTable.UpdateStatus(testNonExistTodoID, todos.Status_DONE.String())
	assert.Nil(t, todo)
	assert.NotNil(t, err)
}

func TestUpdateTodoTypeDone(t *testing.T) {
	todo, err := todosTable.UpdateStatus(testTodoID, todos.Status_DONE.String())
	assert.NotNil(t, todo)
	assert.Nil(t, err)
	assert.Equal(t, todos.Status_DONE.String(), todo.Status)
	assert.True(t, todo.DoneAt.Unix() > 0)
}

func TestUpdateTodoTypeTodo(t *testing.T) {
	todo, err := todosTable.UpdateStatus(testTodoID, todos.Status_TODO.String())
	assert.NotNil(t, todo)
	assert.Nil(t, err)
	assert.Equal(t, todos.Status_TODO.String(), todo.Status)
	assert.Zero(t, todo.DoneAt.Unix())
}

func TestUpdateOrderFail(t *testing.T) {
	todo, err := todosTable.UpdateOrder(testNonExistTodoID, testUpdatedOrder)
	assert.Nil(t, todo)
	assert.NotNil(t, err)
}

func TestUpdateOrder(t *testing.T) {
	todo, err := todosTable.UpdateOrder(testTodoID, testUpdatedOrder)
	assert.NotNil(t, todo)
	assert.Nil(t, err)
	assert.Equal(t, testUpdatedOrder, todo.Order)
}

func TestDeleteByIDFail(t *testing.T) {
	err := todosTable.DeleteByID(testNonExistTodoID)
	assert.NotNil(t, err)
}

func TestDeleteByID(t *testing.T) {
	err := todosTable.DeleteByID(testTodoID)
	assert.Nil(t, err)
}
