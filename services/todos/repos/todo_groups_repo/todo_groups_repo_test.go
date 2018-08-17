package todo_groups_repo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taeho-io/family/services/todos/models"
)

var (
	todoGroupsTable *Table

	testFullTableName       = "family-development-todos-todo_groups"
	testTodoGroupID         = "test_todo_group_id"
	testNonExistTodoGroupID = "test_non_exist_todo_group_id"
	testTitle               = "test_title"
)

func TestMain(m *testing.M) {
	todoGroupsTable = NewMock()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestFullTableName(t *testing.T) {
	assert.Equal(t, testFullTableName, todoGroupsTable.Table().Name())
}

func TestValidationTodoGroupInput(t *testing.T) {
	todoGroup := &models.TodoGroup{
		TodoGroupID: testTodoGroupID,
		Title:       testTitle,
	}
	err := todoGroupsTable.validateTodoGroupInput(todoGroup)
	assert.Nil(t, err)
}

func TestValidationTodoGroupInputInvalidTodoGroupID(t *testing.T) {
	todoGroup := &models.TodoGroup{
		Title: testTitle,
	}
	err := todoGroupsTable.validateTodoGroupInput(todoGroup)
	assert.Equal(t, InvalidTodoGroupIDError, err)
}

func TestValidationTodoGroupInputInvalidTitle(t *testing.T) {
	todoGroup := &models.TodoGroup{
		TodoGroupID: testTodoGroupID,
	}
	err := todoGroupsTable.validateTodoGroupInput(todoGroup)
	assert.Equal(t, InvalidTitleError, err)
}

func TestPutFail(t *testing.T) {
	todoGroup := &models.TodoGroup{}
	err := todoGroupsTable.Put(todoGroup)
	assert.NotNil(t, err)
}

func TestPut(t *testing.T) {
	todoGroup := &models.TodoGroup{
		TodoGroupID: testTodoGroupID,
		Title:       testTitle,
	}
	err := todoGroupsTable.Put(todoGroup)
	assert.Nil(t, err)
}

func TestGetByIDFail(t *testing.T) {
	todoGroup, err := todoGroupsTable.GetByID(testNonExistTodoGroupID)
	assert.Nil(t, todoGroup)
	assert.NotNil(t, err)
}

func TestGetByID(t *testing.T) {
	todoGroup, err := todoGroupsTable.GetByID(testTodoGroupID)
	assert.NotNil(t, todoGroup)
	assert.Nil(t, err)
	assert.Equal(t, testTodoGroupID, todoGroup.TodoGroupID)
}

func TestListByIDsFail(t *testing.T) {
	todoGroupIDs := []string{testNonExistTodoGroupID}
	todoGroupList, err := todoGroupsTable.ListByIDs(todoGroupIDs)
	assert.Nil(t, todoGroupList)
	assert.NotNil(t, err)
}

func TestListByIDs(t *testing.T) {
	todoGroupIDs := []string{testTodoGroupID}
	todoGroupList, err := todoGroupsTable.ListByIDs(todoGroupIDs)
	assert.NotNil(t, todoGroupList)
	assert.Nil(t, err)
	assert.Len(t, todoGroupList, len(todoGroupIDs))
}

func TestDeleteByID(t *testing.T) {
	err := todoGroupsTable.DeleteByID(testTodoGroupID)
	assert.Nil(t, err)
}
