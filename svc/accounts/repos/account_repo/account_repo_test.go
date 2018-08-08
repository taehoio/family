package account_repo

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/taeho-io/family/svc/accounts/models"
)

const (
	testAccountID         = "test_account_id"
	testAccountType       = "email"
	testEmail             = "test@test.io"
	testHashedPassword    = "hashed_password"
	testNewHashedPassword = "new_hashed_password"
	testFullName          = "John Doe"
	testNewFullName       = "Jane Doe"
)

var (
	accountTable *Table
)

func TestMain(m *testing.M) {
	accountTable = NewMock()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestFullTableName(t *testing.T) {
	assert.Equal(t, "family-development-accounts-accounts", accountTable.Table().Name())
}

func TestPut(t *testing.T) {
	account := &models.Account{
		AccountID:      testAccountID,
		Type:           testAccountType,
		Email:          testEmail,
		HashedPassword: testHashedPassword,
		FullName:       testFullName,
		CreateAt:       time.Now(),
		UpdatedAt:      time.Now(),
	}
	err := accountTable.Put(account)
	assert.Nil(t, err)
}

func TestGet(t *testing.T) {
	account, err := accountTable.Get(testAccountID)
	assert.Nil(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, testEmail, account.Email)
}

func TestGetByEmail(t *testing.T) {
	account, err := accountTable.GetByEmail(testEmail)
	assert.Nil(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, testAccountID, account.AccountID)
}

func TestUpdateHashedPassword(t *testing.T) {
	updatedAccount, err := accountTable.UpdateHashedPassword(testAccountID, testNewHashedPassword)
	assert.Nil(t, err)
	assert.NotNil(t, updatedAccount)
	assert.Equal(t, testNewHashedPassword, updatedAccount.HashedPassword)
}

func TestUpdateFullName(t *testing.T) {
	updatedAccount, err := accountTable.UpdateFullName(testAccountID, testNewFullName)
	assert.Nil(t, err)
	assert.NotNil(t, updatedAccount)
	assert.Equal(t, testNewFullName, updatedAccount.FullName)
}

func TestDelete(t *testing.T) {
	err := accountTable.Delete(testAccountID)
	assert.Nil(t, err)
}
