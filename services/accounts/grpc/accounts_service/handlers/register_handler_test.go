package handlers

import (
	"testing"

	"github.com/guregu/dynamo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
	"github.com/taeho-io/family/services/accounts/crypt"
	mockCrypt "github.com/taeho-io/family/services/accounts/mocks/crypt"
	mockAccountsRepoMock "github.com/taeho-io/family/services/accounts/mocks/repos/accounts_repo"
	"github.com/taeho-io/family/services/accounts/repos/accounts_repo"
)

func TestRegisterInvalidAuthTypeError(t *testing.T) {
	accountsTableMock := accounts_repo.NewMock()
	cryptMock := crypt.New()
	register := Register(accountsTableMock, cryptMock)

	ctx := context.Background()
	req := &accounts.RegisterRequest{
		AuthType: accounts.AuthType_NONE,
		FullName: testFullName,
		Password: testPassword,
	}
	res, err := register(ctx, req)
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, InvalidAuthTypeError, err)
}

func TestRegisterInvalidEmailError(t *testing.T) {
	accountsTableMock := accounts_repo.NewMock()
	cryptMock := crypt.New()
	register := Register(accountsTableMock, cryptMock)

	ctx := context.Background()
	req := &accounts.RegisterRequest{
		AuthType: testAuthTypeEmail,
		FullName: testFullName,
		Password: testPassword,
	}
	res, err := register(ctx, req)
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, InvalidEmailError, err)
}

func TestRegisterInvalidFullNameError(t *testing.T) {
	accountsTableMock := accounts_repo.NewMock()
	cryptMock := crypt.New()
	register := Register(accountsTableMock, cryptMock)

	ctx := context.Background()
	req := &accounts.RegisterRequest{
		AuthType: testAuthTypeEmail,
		Email:    testEmail,
		Password: testPassword,
	}
	res, err := register(ctx, req)
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, InvalidFullNameError, err)
}

func TestRegisterInvalidPasswordError(t *testing.T) {
	accountsTableMock := accounts_repo.NewMock()
	cryptMock := crypt.New()
	register := Register(accountsTableMock, cryptMock)

	ctx := context.Background()
	req := &accounts.RegisterRequest{
		AuthType: testAuthTypeEmail,
		FullName: testFullName,
		Email:    testEmail,
	}
	res, err := register(ctx, req)
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, InvalidPasswordError, err)
}

func TestRegisterHashPasswordFail(t *testing.T) {
	accountsTableMock := accounts_repo.NewMock()
	cryptMock := new(mockCrypt.IFace)
	cryptMock.On("HashPassword", testPassword).
		Return("", bcrypt.ErrHashTooShort)
	register := Register(accountsTableMock, cryptMock)

	ctx := context.Background()
	req := &accounts.RegisterRequest{
		AuthType: testAuthTypeEmail,
		FullName: testFullName,
		Email:    testEmail,
		Password: testPassword,
	}
	res, err := register(ctx, req)
	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, bcrypt.ErrHashTooShort, err)
}

func TestRegisterPutError(t *testing.T) {
	accountsTableMock := new(mockAccountsRepoMock.IFace)
	accountsTableMock.On("GetByEmail", testEmail).Return(nil, nil)
	accountsTableMock.On("Put", mock.Anything).Return(dynamo.ErrTooMany)
	cryptMock := crypt.New()
	register := Register(accountsTableMock, cryptMock)

	ctx := context.Background()
	req := &accounts.RegisterRequest{
		AuthType: testAuthTypeEmail,
		FullName: testFullName,
		Email:    testEmail,
		Password: testPassword,
	}
	res, err := register(ctx, req)
	assert.Nil(t, res)
	assert.NotNil(t, err)
}

func TestRegister(t *testing.T) {
	accountsTableMock := accounts_repo.NewMock()
	cryptMock := crypt.New()
	register := Register(accountsTableMock, cryptMock)

	ctx := context.Background()
	req := &accounts.RegisterRequest{
		AuthType: testAuthTypeEmail,
		FullName: testFullName,
		Email:    testEmail,
		Password: testPassword,
	}
	res, err := register(ctx, req)
	assert.NotNil(t, res)
	assert.Nil(t, err)
	registeredAccountID := res.AccountId

	err = accountsTableMock.DeleteByID(registeredAccountID)
	assert.Nil(t, err)
}

func TestRegisterEmailAlreadyExistsError(t *testing.T) {
	accountsTableMock := accounts_repo.NewMock()
	cryptMock := crypt.New()
	register := Register(accountsTableMock, cryptMock)

	ctx := context.Background()
	req := &accounts.RegisterRequest{
		AuthType: testAuthTypeEmail,
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

	err = accountsTableMock.DeleteByID(registeredAccountID)
	assert.Nil(t, err)
}
