package handler

import (
	"testing"

	"github.com/guregu/dynamo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
	"github.com/taeho-io/family/services/accounts/internal/repo"
	"github.com/taeho-io/family/services/accounts/mocks"
	"github.com/taeho-io/family/services/accounts/pkg/crypt"
)

func TestRegisterInvalidAuthTypeError(t *testing.T) {
	accountsRepoMock := repo.NewMockAccountsRepo()
	cryptMock := crypt.NewMock()
	register := Register(accountsRepoMock, cryptMock)

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
	accountsRepoMock := repo.NewMockAccountsRepo()
	cryptMock := crypt.NewMock()
	register := Register(accountsRepoMock, cryptMock)

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
	accountsRepoMock := repo.NewMockAccountsRepo()
	cryptMock := crypt.NewMock()
	register := Register(accountsRepoMock, cryptMock)

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
	accountsRepoMock := repo.NewMockAccountsRepo()
	cryptMock := crypt.NewMock()
	register := Register(accountsRepoMock, cryptMock)

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
	accountsRepoMock := repo.NewMockAccountsRepo()
	cryptMock := new(mocks.Crypt)
	cryptMock.On("HashPassword", testPassword).
		Return("", bcrypt.ErrHashTooShort)
	register := Register(accountsRepoMock, cryptMock)

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
	accountsRepoMock := new(mocks.AccountsRepo)
	accountsRepoMock.On("GetByEmail", testEmail).Return(nil, nil)
	accountsRepoMock.On("Put", mock.Anything).Return(dynamo.ErrTooMany)
	cryptMock := crypt.NewMock()
	register := Register(accountsRepoMock, cryptMock)

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
	accountsRepoMock := repo.NewMockAccountsRepo()
	cryptMock := crypt.NewMock()
	register := Register(accountsRepoMock, cryptMock)

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

	err = accountsRepoMock.DeleteByID(registeredAccountID)
	assert.Nil(t, err)
}

func TestRegisterEmailAlreadyExistsError(t *testing.T) {
	accountsRepoMock := repo.NewMockAccountsRepo()
	cryptMock := crypt.NewMock()
	register := Register(accountsRepoMock, cryptMock)

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

	err = accountsRepoMock.DeleteByID(registeredAccountID)
	assert.Nil(t, err)
}
