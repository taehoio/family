package repo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todogroups"
	"github.com/taeho-io/family/services/todogroups/internal/model"
)

var (
	todoGroupPermitsRepo TodoGroupPermitsRepo

	testAccountID                 = "test_account_id"
	testNonExistAccountID         = "test_non_exist_account_id"
	testAnotherTodoGroupID        = "test_another_todo_group_id"
	testTodoGroupPermitTypeOwner  = todogroups.PermitType_PERMIT_TYPE_OWNER
	testTodoGroupPermitTypeEditor = todogroups.PermitType_PERMIT_TYPE_EDITOR
	testWongTodoGroupPermitType   = todogroups.PermitType(999)
)

func TestValidationTodoGroupPermitInput(t *testing.T) {
	todoGroupPermit := &model.TodoGroupPermit{
		AccountID:   testAccountID,
		TodoGroupID: testTodoGroupID,
		PermitType:  testTodoGroupPermitTypeOwner,
	}
	err := validateTodoGroupPermitInput(todoGroupPermit)
	assert.Nil(t, err)
}

func TestValidateTodoGroupPermitInputInvalidAccountID(t *testing.T) {
	todoGroupPermit := &model.TodoGroupPermit{
		TodoGroupID: testTodoGroupID,
		PermitType:  testTodoGroupPermitTypeOwner,
	}
	err := validateTodoGroupPermitInput(todoGroupPermit)
	assert.Equal(t, InvalidAccountIDError, err)
}

func TestValidateTodoGroupPermitInputInvalidTodoGroupID(t *testing.T) {
	todoGroupPermit := &model.TodoGroupPermit{
		AccountID:  testAccountID,
		PermitType: testTodoGroupPermitTypeOwner,
	}
	err := validateTodoGroupPermitInput(todoGroupPermit)
	assert.Equal(t, InvalidTodoGroupIDError, err)
}

func TestValidateTodoGroupPermitInputInvalidPermitType(t *testing.T) {
	todoGroupPermit := &model.TodoGroupPermit{
		AccountID:   testAccountID,
		TodoGroupID: testTodoGroupID,
		PermitType:  testWongTodoGroupPermitType,
	}
	err := validateTodoGroupPermitInput(todoGroupPermit)
	assert.Equal(t, InvalidPermitTypeError, err)
}

func TestPutFail(t *testing.T) {
	todoGroupPermit := &model.TodoGroupPermit{}
	err := todoGroupPermitsRepo.Put(todoGroupPermit)
	assert.NotNil(t, err)
}

func TestPut(t *testing.T) {
	fmt.Println("!!!")
	fmt.Println(todoGroupPermitsRepo.Table().Name())

	todoGroupPermit := &model.TodoGroupPermit{
		AccountID:   testAccountID,
		TodoGroupID: testTodoGroupID,
		PermitType:  testTodoGroupPermitTypeOwner,
	}
	err := todoGroupPermitsRepo.Put(todoGroupPermit)
	assert.Nil(t, err)
}

func TestPutOneMore(t *testing.T) {
	todoGroupPermit := &model.TodoGroupPermit{
		AccountID:   testAccountID,
		TodoGroupID: testAnotherTodoGroupID,
		PermitType:  testTodoGroupPermitTypeEditor,
	}
	err := todoGroupPermitsRepo.Put(todoGroupPermit)
	assert.Nil(t, err)
}

func TestGetFail(t *testing.T) {
	todoGroupPermit, err := todoGroupPermitsRepo.Get(testNonExistAccountID, testTodoGroupID)
	assert.Nil(t, todoGroupPermit)
	assert.NotNil(t, err)
}

func TestGet(t *testing.T) {
	todoGroupPermit, err := todoGroupPermitsRepo.Get(testAccountID, testTodoGroupID)
	assert.NotNil(t, todoGroupPermit)
	assert.Nil(t, err)
}

func TestListByAccountID(t *testing.T) {
	todoGroupPermits, err := todoGroupPermitsRepo.ListByAccountID(testAccountID)
	assert.NotNil(t, todoGroupPermits)
	assert.Len(t, todoGroupPermits, 2)
	assert.Nil(t, err)
}

func TestListByAccountIDEmpty(t *testing.T) {
	todoGroupPermits, err := todoGroupPermitsRepo.ListByAccountID(testNonExistAccountID)
	assert.Len(t, todoGroupPermits, 0)
	assert.Nil(t, err)
}

func TestDeleteFail(t *testing.T) {
	err := todoGroupPermitsRepo.Delete(testNonExistAccountID, testTodoGroupID)
	assert.NotNil(t, err)
}

func TestDelete(t *testing.T) {
	err := todoGroupPermitsRepo.Delete(testAccountID, testTodoGroupID)
	assert.Nil(t, err)
}

func TestListByAccountIDFail(t *testing.T) {
	accountIDFieldKey = "wrong_account_id"
	todoGroupPermits, err := todoGroupPermitsRepo.ListByAccountID(testNonExistAccountID)
	assert.Nil(t, todoGroupPermits)
	assert.NotNil(t, err)
}
