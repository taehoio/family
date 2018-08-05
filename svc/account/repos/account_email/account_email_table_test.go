package account_email

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testEmail     = "test@test.io"
	testAccountID = "test_account_id"
)

func TestFullTableName(t *testing.T) {
	accountEmailTable := NewMock()

	assert.Equal(t, "family-testing-account-account_email", accountEmailTable.Table().Name())
}

func TestPut(t *testing.T) {
	accountEmailTable := NewMock()

	accountEmail := &AccountEmail{
		Email:     testEmail,
		AccountID: testAccountID,
	}
	err := accountEmailTable.Put(accountEmail)
	assert.Nil(t, err)
}

func TestGet(t *testing.T) {
	accountEmailTable := NewMock()

	accountEmail, err := accountEmailTable.Get(testEmail)
	assert.Nil(t, err)
	assert.NotNil(t, accountEmail)
}

func TestGetAccountIDByEmail(t *testing.T) {
	accountEmailTable := NewMock()

	accountID, err := accountEmailTable.GetAccountIDByEmail(testEmail)
	assert.Nil(t, err)
	assert.Equal(t, testAccountID, accountID)
}

func TestDelete(t *testing.T) {
	accountEmailTable := NewMock()

	err := accountEmailTable.Delete(testEmail)
	assert.Nil(t, err)
}
