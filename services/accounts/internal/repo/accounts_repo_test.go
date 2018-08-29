package repo

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
	"github.com/taeho-io/family/services/accounts/internal/model"
)

const (
	testFullTableName     = "family-development-accounts-accounts"
	testAccountID         = "test_account_id"
	testNonExistAccountID = "test_non_exist_account_id"
	testAccountType       = accounts.AuthType_EMAIL
	testEmail             = "test@test.io"
	testHashedPassword    = "hashed_password"
	testNewHashedPassword = "new_hashed_password"
	testFullName          = "John Doe"
	testNewFullName       = "Jane Doe"
)

var (
	accountsRepo AccountsRepo
)

func TestMain(m *testing.M) {
	accountsRepo = NewMockAccountsRepo()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestFullTableName(t *testing.T) {
	assert.Equal(t, testFullTableName, accountsRepo.Table().Name())
}

func TestPut(t *testing.T) {
	account := &model.Account{
		AccountID:      testAccountID,
		Type:           testAccountType,
		Email:          testEmail,
		HashedPassword: testHashedPassword,
		FullName:       testFullName,
		CreateAt:       time.Now(),
		UpdatedAt:      time.Now(),
	}
	err := accountsRepo.Put(account)
	assert.Nil(t, err)
}

func TestGetFail(t *testing.T) {
	account, err := accountsRepo.GetByID(testNonExistAccountID)
	assert.Nil(t, account)
	assert.NotNil(t, err)
}

func TestGet(t *testing.T) {
	account, err := accountsRepo.GetByID(testAccountID)
	assert.NotNil(t, account)
	assert.Nil(t, err)
	assert.Equal(t, testEmail, account.Email)
}

func TestGetByEmailFail(t *testing.T) {
	account, err := accountsRepo.GetByEmail(testNonExistAccountID)
	assert.Nil(t, account)
	assert.NotNil(t, err)
}

func TestGetByEmail(t *testing.T) {
	account, err := accountsRepo.GetByEmail(testEmail)
	assert.NotNil(t, account)
	assert.Nil(t, err)
	assert.Equal(t, testAccountID, account.AccountID)
}

func TestUpdateHashedPasswordFail(t *testing.T) {
	updatedAccount, err := accountsRepo.UpdateHashedPassword(testNonExistAccountID, testNewHashedPassword)
	assert.Nil(t, updatedAccount)
	assert.NotNil(t, err)
}

func TestUpdateHashedPassword(t *testing.T) {
	updatedAccount, err := accountsRepo.UpdateHashedPassword(testAccountID, testNewHashedPassword)
	assert.NotNil(t, updatedAccount)
	assert.Nil(t, err)
	assert.Equal(t, testNewHashedPassword, updatedAccount.HashedPassword)
}

func TestUpdateFullNameFail(t *testing.T) {
	updatedAccount, err := accountsRepo.UpdateFullName(testNonExistAccountID, testNewFullName)
	assert.Nil(t, updatedAccount)
	assert.NotNil(t, err)
}

func TestUpdateFullName(t *testing.T) {
	updatedAccount, err := accountsRepo.UpdateFullName(testAccountID, testNewFullName)
	assert.NotNil(t, updatedAccount)
	assert.Nil(t, err)
	assert.Equal(t, testNewFullName, updatedAccount.FullName)
}

func TestDelete(t *testing.T) {
	err := accountsRepo.DeleteByID(testAccountID)
	assert.Nil(t, err)
}
