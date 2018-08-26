package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
)

func TestLogIn(t *testing.T) {
	ctx := context.Background()
	registerReq := &accounts.RegisterRequest{
		FullName: testFullName,
		Email:    testEmail,
		Password: testPassword,
	}
	registerRes, err := register(ctx, registerReq)
	assert.NotNil(t, registerRes)
	assert.Nil(t, err)
	registeredAccountID := registerRes.AccountId

	logInReq := &accounts.LogInRequest{
		AuthType: testAuthTypeEmail,
		Email:    testEmail,
		Password: testPassword,
	}
	logInRes, err := logIn(ctx, logInReq)
	assert.NotNil(t, logInRes)
	assert.Nil(t, err)

	err = accountsTableMock.Delete(registeredAccountID)
	assert.Nil(t, err)
}
