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
	testAnotherTodoID      = "test_another_todo_id"
	testNonExistTodoID     = "test_no_exist_todo_id"
	testParentType         = todos.ParentType_PARENT_TYPE_TODO_GROUP
	testParentID           = "test_parent_id"
	testTodoTypeTodo       = todos.Status_STATUS_DONE
	testTitle              = "test_title"
	testDescription        = "test_description"
	testUpdatedParentType  = todos.ParentType_PARENT_TYPE_TODO_GROUP
	testUpdatedParentID    = "test_updated_parent_id"
	testUpdatedTitle       = "test_updated_title"
	testUpdatedDescription = "test_updated_description"
	testUpdatedOrder       = "m"
	testUpdatedAssignedTo  = "test_assigned_to"
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
		ParentType:  testParentType,
		ParentID:    testParentID,
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
		ParentType:  testParentType,
		ParentID:    testParentID,
		Title:       testTitle,
		Description: testDescription,
		Status:      testTodoTypeTodo,
	}
	err := todosTable.validateTodoInput(todo)
	assert.Equal(t, InvalidTodoIDError, err)
}

func TestValidateTodoInputInvalidParentID(t *testing.T) {
	todo := &models.Todo{
		TodoID:      testTodoID,
		ParentType:  testParentType,
		Title:       testTitle,
		Description: testDescription,
		Status:      testTodoTypeTodo,
	}
	err := todosTable.validateTodoInput(todo)
	assert.Equal(t, InvalidParentIDError, err)
}

func TestValidateTodoInputInvalidTitle(t *testing.T) {
	todo := &models.Todo{
		TodoID:      testTodoID,
		ParentType:  testParentType,
		ParentID:    testParentID,
		Description: testDescription,
		Status:      testTodoTypeTodo,
	}
	err := todosTable.validateTodoInput(todo)
	assert.Equal(t, InvalidTitleError, err)
}

func TestPut(t *testing.T) {
	todo := &models.Todo{
		TodoID:      testTodoID,
		ParentType:  testParentType,
		ParentID:    testParentID,
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
	assert.Equal(t, testParentType, todo.ParentType)
	assert.Equal(t, testParentID, todo.ParentID)
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

func TestPutAnother(t *testing.T) {
	todo := &models.Todo{
		TodoID:      testAnotherTodoID,
		ParentType:  testParentType,
		ParentID:    testParentID,
		Title:       testTitle,
		Description: testDescription,
		Status:      testTodoTypeTodo,
	}
	err := todosTable.Put(todo)
	assert.Nil(t, err)
}

func TestListByIDs(t *testing.T) {
	todoIDs := []string{testTodoID, testAnotherTodoID}
	todoList, err := todosTable.ListByIDs(todoIDs)
	assert.Nil(t, err)
	assert.NotNil(t, todoList)
	assert.Len(t, todoList, 2)

	todo := todoList[0]
	assert.Equal(t, testTodoID, todo.TodoID)
	assert.Equal(t, testParentType, todo.ParentType)
	assert.Equal(t, testParentID, todo.ParentID)
	assert.Equal(t, testTitle, todo.Title)

	todo = todoList[1]
	assert.Equal(t, testAnotherTodoID, todo.TodoID)
	assert.Equal(t, testParentType, todo.ParentType)
	assert.Equal(t, testParentID, todo.ParentID)
	assert.Equal(t, testTitle, todo.Title)
}

func TestListByParentID(t *testing.T) {
	todoList, err := todosTable.ListByParentID(testParentID)
	assert.Nil(t, err)
	assert.NotNil(t, todoList)
	assert.Len(t, todoList, 2)

	todo := todoList[0]
	assert.Equal(t, testTodoID, todo.TodoID)
	assert.Equal(t, testParentType, todo.ParentType)
	assert.Equal(t, testParentID, todo.ParentID)
	assert.Equal(t, testTitle, todo.Title)

	todo = todoList[1]
	assert.Equal(t, testAnotherTodoID, todo.TodoID)
	assert.Equal(t, testParentType, todo.ParentType)
	assert.Equal(t, testParentID, todo.ParentID)
	assert.Equal(t, testTitle, todo.Title)
}

func TestUpdateParentFail(t *testing.T) {
	todo, err := todosTable.UpdateParent(testNonExistTodoID, testUpdatedParentType, testUpdatedParentID)
	assert.Nil(t, todo)
	assert.NotNil(t, err)
}

func TestUpdateParent(t *testing.T) {
	todo, err := todosTable.UpdateParent(testTodoID, testUpdatedParentType, testUpdatedParentID)
	assert.NotNil(t, todo)
	assert.Nil(t, err)
	assert.Equal(t, testUpdatedParentType, todo.ParentType)
	assert.Equal(t, testUpdatedParentID, todo.ParentID)
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
	todo, err := todosTable.UpdateStatus(testNonExistTodoID, todos.Status_STATUS_DONE)
	assert.Nil(t, todo)
	assert.NotNil(t, err)
}

func TestUpdateTodoTypeDone(t *testing.T) {
	todo, err := todosTable.UpdateStatus(testTodoID, todos.Status_STATUS_DONE)
	assert.NotNil(t, todo)
	assert.Nil(t, err)
	assert.Equal(t, todos.Status_STATUS_DONE, todo.Status)
	assert.True(t, todo.DoneAt.Unix() > 0)
}

func TestUpdateTodoTypeTodo(t *testing.T) {
	todo, err := todosTable.UpdateStatus(testTodoID, todos.Status_STATUS_TODO)
	assert.NotNil(t, todo)
	assert.Nil(t, err)
	assert.Equal(t, todos.Status_STATUS_TODO, todo.Status)
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

func TestUpdateAssignedToFail(t *testing.T) {
	todo, err := todosTable.UpdateAssignedTo(testNonExistTodoID, testUpdatedAssignedTo)
	assert.Nil(t, todo)
	assert.NotNil(t, err)
}

func TestUpdateAssignedTo(t *testing.T) {
	todo, err := todosTable.UpdateAssignedTo(testTodoID, testUpdatedAssignedTo)
	assert.NotNil(t, todo)
	assert.Nil(t, err)
	assert.Equal(t, testUpdatedAssignedTo, todo.AssignedTo)
}

func TestDeleteByIDFail(t *testing.T) {
	err := todosTable.DeleteByID(testNonExistTodoID)
	assert.NotNil(t, err)
}

func TestDeleteByID(t *testing.T) {
	err := todosTable.DeleteByID(testTodoID)
	assert.Nil(t, err)
}

func TestListByParentIDFail(t *testing.T) {
	parentIDFieldKey = "wrong_parent_id"
	todoList, err := todosTable.ListByParentID(testParentID)
	assert.NotNil(t, err)
	assert.Nil(t, todoList)
}
