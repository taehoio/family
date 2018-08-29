package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/taeho-io/family/services/todogroups/internal/model"
)

var (
	todoGroupsRepo GroupsRepo

	testTodoGroupID         = "test_todo_group_id"
	testNonExistTodoGroupID = "test_non_exist_todo_group_id"
	testTitle               = "test_title"
	testUpdatedTitle        = "test_updated_title"
	testUpdatedDescription  = "test_updated_description"
	testUpdatedOrder        = "test_updated_order"
	testCreatedBy           = "test_created_by"
)

func TestValidateTodoGroupInput(t *testing.T) {
	todoGroup := &model.TodoGroup{
		TodoGroupID: testTodoGroupID,
		Title:       testTitle,
		CreatedBy:   testCreatedBy,
	}
	err := validateTodoGroupInput(todoGroup)
	assert.Nil(t, err)
}

func TestValidateTodoGroupInputInvalidTodoGroupID(t *testing.T) {
	todoGroup := &model.TodoGroup{
		Title: testTitle,
	}
	err := validateTodoGroupInput(todoGroup)
	assert.Equal(t, InvalidTodoGroupIDError, err)
}

func TestValidateTodoGroupInputInvalidTitle(t *testing.T) {
	todoGroup := &model.TodoGroup{
		TodoGroupID: testTodoGroupID,
	}
	err := validateTodoGroupInput(todoGroup)
	assert.Equal(t, InvalidTitleError, err)
}

func TestValidateTodoGroupInputInvalidCreatedBy(t *testing.T) {
	todoGroup := &model.TodoGroup{
		Title:       testTitle,
		TodoGroupID: testTodoGroupID,
	}
	err := validateTodoGroupInput(todoGroup)
	assert.Equal(t, InvalidCreatedByError, err)
}

func TestTodoGroupsRepoPutFail(t *testing.T) {
	todoGroup := &model.TodoGroup{}
	err := todoGroupsRepo.Put(todoGroup)
	assert.NotNil(t, err)
}

func TestTodoGroupsRepoPut(t *testing.T) {
	todoGroup := &model.TodoGroup{
		TodoGroupID: testTodoGroupID,
		Title:       testTitle,
		CreatedBy:   testCreatedBy,
	}
	err := todoGroupsRepo.Put(todoGroup)
	assert.Nil(t, err)
}

func TestGetByIDFail(t *testing.T) {
	todoGroup, err := todoGroupsRepo.GetByID(testNonExistTodoGroupID)
	assert.Nil(t, todoGroup)
	assert.NotNil(t, err)
}

func TestGetByID(t *testing.T) {
	todoGroup, err := todoGroupsRepo.GetByID(testTodoGroupID)
	assert.NotNil(t, todoGroup)
	assert.Nil(t, err)
	assert.Equal(t, testTodoGroupID, todoGroup.TodoGroupID)
}

func TestListByIDsFail(t *testing.T) {
	todoGroupIDs := []string{testNonExistTodoGroupID}
	todoGroupList, err := todoGroupsRepo.ListByIDs(todoGroupIDs)
	assert.Nil(t, todoGroupList)
	assert.NotNil(t, err)
}

func TestListByIDs(t *testing.T) {
	todoGroupIDs := []string{testTodoGroupID}
	todoGroupList, err := todoGroupsRepo.ListByIDs(todoGroupIDs)
	assert.NotNil(t, todoGroupList)
	assert.Nil(t, err)
	assert.Len(t, todoGroupList, len(todoGroupIDs))
}

func TestUpdateTitleFail(t *testing.T) {
	todoGroup, err := todoGroupsRepo.UpdateTitle(testNonExistTodoGroupID, testUpdatedTitle)
	assert.Nil(t, todoGroup)
	assert.NotNil(t, err)
}

func TestUpdateTitle(t *testing.T) {
	todoGroup, err := todoGroupsRepo.UpdateTitle(testTodoGroupID, testUpdatedTitle)
	assert.NotNil(t, todoGroup)
	assert.Nil(t, err)
	assert.Equal(t, testUpdatedTitle, todoGroup.Title)
}

func TestUpdateDescriptionFail(t *testing.T) {
	todoGroup, err := todoGroupsRepo.UpdateDescription(testNonExistTodoGroupID, testUpdatedDescription)
	assert.Nil(t, todoGroup)
	assert.NotNil(t, err)
}

func TestUpdateDescription(t *testing.T) {
	todoGroup, err := todoGroupsRepo.UpdateDescription(testTodoGroupID, testUpdatedDescription)
	assert.NotNil(t, todoGroup)
	assert.Nil(t, err)
	assert.Equal(t, testUpdatedDescription, todoGroup.Description)
}

func TestUpdateOrderFail(t *testing.T) {
	todoGroup, err := todoGroupsRepo.UpdateOrder(testNonExistTodoGroupID, testUpdatedOrder)
	assert.Nil(t, todoGroup)
	assert.NotNil(t, err)
}

func TestUpdateOrder(t *testing.T) {
	todoGroup, err := todoGroupsRepo.UpdateOrder(testTodoGroupID, testUpdatedOrder)
	assert.NotNil(t, todoGroup)
	assert.Nil(t, err)
	assert.Equal(t, testUpdatedDescription, todoGroup.Description)
}

func TestDeleteByID(t *testing.T) {
	err := todoGroupsRepo.DeleteByID(testTodoGroupID)
	assert.Nil(t, err)
}
