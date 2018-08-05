package account_email_repo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taeho-io/family/svc/accounts/models"
)

const (
	testEmail     = "test@test.io"
	testAccountID = "test_account_id"
)

var (
	accountEmailTable *Table
)

func TestMain(m *testing.M) {
	accountEmailTable = NewMock()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestFullTableName(t *testing.T) {
	assert.Equal(t, "family-development-accounts-accounts_emails", accountEmailTable.Table().Name())
}

func TestPut(t *testing.T) {
	accountEmail := &models.AccountEmail{
		Email:     testEmail,
		AccountID: testAccountID,
	}
	err := accountEmailTable.Put(accountEmail)
	assert.Nil(t, err)
}

func TestGet(t *testing.T) {
	accountEmail, err := accountEmailTable.Get(testEmail)
	assert.Nil(t, err)
	assert.NotNil(t, accountEmail)
	assert.Equal(t, testAccountID, accountEmail.AccountID)
}

func TestGetAccountIDByEmail(t *testing.T) {
	accountID, err := accountEmailTable.GetAccountIDByEmail(testEmail)
	assert.Nil(t, err)
	assert.Equal(t, testAccountID, accountID)
}

func TestDelete(t *testing.T) {
	accountEmailTable := NewMock()

	err := accountEmailTable.Delete(testEmail)
	assert.Nil(t, err)
}
