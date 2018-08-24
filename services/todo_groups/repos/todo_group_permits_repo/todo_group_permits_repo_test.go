package todo_group_permits_repo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
	"github.com/taeho-io/family/services/todo_groups/models"
)

var (
	todoGroupPermitsTable *Table

	testFullTableName             = "family-development-todo_groups-todo_group_permits"
	testAccountID                 = "test_account_id"
	testNonExistAccountID         = "test_non_exist_account_id"
	testTodoGroupID               = "test_todo_group_id"
	testAnotherTodoGroupID        = "test_another_todo_group_id"
	testTodoGroupPermitTypeOwner  = todo_groups.TodoGroupPermitType_OWNER.String()
	testTodoGroupPermitTypeEditor = todo_groups.TodoGroupPermitType_EDITOR.String()
)

func TestMain(m *testing.M) {
	todoGroupPermitsTable = NewMock()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestFullTableName(t *testing.T) {
	assert.Equal(t, testFullTableName, todoGroupPermitsTable.Table().Name())
}

func TestValidationTodoGroupPermitInput(t *testing.T) {
	todoGroupPermit := &models.TodoGroupPermit{
		AccountID:   testAccountID,
		TodoGroupID: testTodoGroupID,
		PermitType:  testTodoGroupPermitTypeOwner,
	}
	err := todoGroupPermitsTable.validateTodoGroupPermitInput(todoGroupPermit)
	assert.Nil(t, err)
}

func TestValidateTodoGroupPermitInputInvalidAccountID(t *testing.T) {
	todoGroupPermit := &models.TodoGroupPermit{
		TodoGroupID: testTodoGroupID,
		PermitType:  testTodoGroupPermitTypeOwner,
	}
	err := todoGroupPermitsTable.validateTodoGroupPermitInput(todoGroupPermit)
	assert.Equal(t, InvalidAccountIDError, err)
}

func TestValidateTodoGroupPermitInputInvalidTodoGroupID(t *testing.T) {
	todoGroupPermit := &models.TodoGroupPermit{
		AccountID:  testAccountID,
		PermitType: testTodoGroupPermitTypeOwner,
	}
	err := todoGroupPermitsTable.validateTodoGroupPermitInput(todoGroupPermit)
	assert.Equal(t, InvalidTodoGroupIDError, err)
}

func TestValidateTodoGroupPermitInputInvalidPermitType(t *testing.T) {
	todoGroupPermit := &models.TodoGroupPermit{
		AccountID:   testAccountID,
		TodoGroupID: testTodoGroupID,
	}
	err := todoGroupPermitsTable.validateTodoGroupPermitInput(todoGroupPermit)
	assert.Equal(t, InvalidPermitTypeError, err)
}

func TestPutFail(t *testing.T) {
	todoGroupPermit := &models.TodoGroupPermit{}
	err := todoGroupPermitsTable.Put(todoGroupPermit)
	assert.NotNil(t, err)
}

func TestPut(t *testing.T) {
	todoGroupPermit := &models.TodoGroupPermit{
		AccountID:   testAccountID,
		TodoGroupID: testTodoGroupID,
		PermitType:  testTodoGroupPermitTypeOwner,
	}
	err := todoGroupPermitsTable.Put(todoGroupPermit)
	assert.Nil(t, err)
}

func TestPutOneMore(t *testing.T) {
	todoGroupPermit := &models.TodoGroupPermit{
		AccountID:   testAccountID,
		TodoGroupID: testAnotherTodoGroupID,
		PermitType:  testTodoGroupPermitTypeEditor,
	}
	err := todoGroupPermitsTable.Put(todoGroupPermit)
	assert.Nil(t, err)
}

func TestGetFail(t *testing.T) {
	todoGroupPermit, err := todoGroupPermitsTable.Get(testNonExistAccountID, testTodoGroupID)
	assert.Nil(t, todoGroupPermit)
	assert.NotNil(t, err)
}

func TestGet(t *testing.T) {
	todoGroupPermit, err := todoGroupPermitsTable.Get(testAccountID, testTodoGroupID)
	assert.NotNil(t, todoGroupPermit)
	assert.Nil(t, err)
}

func TestListByAccountID(t *testing.T) {
	todoGroupPermits, err := todoGroupPermitsTable.ListByAccountID(testAccountID)
	assert.NotNil(t, todoGroupPermits)
	assert.Len(t, todoGroupPermits, 2)
	assert.Nil(t, err)
}

func TestListByAccountIDEmpty(t *testing.T) {
	todoGroupPermits, err := todoGroupPermitsTable.ListByAccountID(testNonExistAccountID)
	assert.Len(t, todoGroupPermits, 0)
	assert.Nil(t, err)
}

func TestDeleteFail(t *testing.T) {
	err := todoGroupPermitsTable.Delete(testNonExistAccountID, testTodoGroupID)
	assert.NotNil(t, err)
}

func TestDelete(t *testing.T) {
	err := todoGroupPermitsTable.Delete(testAccountID, testTodoGroupID)
	assert.Nil(t, err)
}

func TestListByAccountIDFail(t *testing.T) {
	accountIDFieldKey = "wrong_account_id"
	todoGroupPermits, err := todoGroupPermitsTable.ListByAccountID(testNonExistAccountID)
	assert.Nil(t, todoGroupPermits)
	assert.NotNil(t, err)
}
