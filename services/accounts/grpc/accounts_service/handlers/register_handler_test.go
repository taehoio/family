package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
)

func TestRegisterInvalidEmailError(t *testing.T) {
	ctx := context.Background()
	req := &accounts.RegisterRequest{
		FullName: testFullName,
		Password: testPassword,
	}
	res, err := register(ctx, req)
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, InvalidEmailError, err)
}

func TestRegisterInvalidFullNameError(t *testing.T) {
	ctx := context.Background()
	req := &accounts.RegisterRequest{
		Email:    testEmail,
		Password: testPassword,
	}
	res, err := register(ctx, req)
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, InvalidFullNameError, err)
}

func TestRegisterInvalidPasswordError(t *testing.T) {
	ctx := context.Background()
	req := &accounts.RegisterRequest{
		FullName: testFullName,
		Email:    testEmail,
	}
	res, err := register(ctx, req)
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, InvalidPasswordError, err)
}

func TestRegister(t *testing.T) {
	ctx := context.Background()
	req := &accounts.RegisterRequest{
		FullName: testFullName,
		Email:    testEmail,
		Password: testPassword,
	}
	res, err := register(ctx, req)
	assert.NotNil(t, res)
	assert.Nil(t, err)
	registeredAccountID := res.AccountId

	err = accountsTableMock.Delete(registeredAccountID)
	assert.Nil(t, err)
}

func TestRegisterEmailAlreadyExistsError(t *testing.T) {
	ctx := context.Background()
	req := &accounts.RegisterRequest{
		FullName: testFullName,
		Email:    testEmail,
		Password: testPassword,
	}
	res, err := register(ctx, req)
	assert.NotNil(t, res)
	assert.Nil(t, err)
	registeredAccountID := res.AccountId

	res, err = register(ctx, req)
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, EmailAlreadyExistsError, err)

	err = accountsTableMock.Delete(registeredAccountID)
	assert.Nil(t, err)
}
