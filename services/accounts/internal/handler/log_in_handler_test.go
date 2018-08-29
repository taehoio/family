package handler

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/services/accounts/internal/model"
	"github.com/taeho-io/family/services/accounts/internal/repo"
	"github.com/taeho-io/family/services/accounts/mocks"
	"github.com/taeho-io/family/services/accounts/pkg/crypt"
)

func TestLogInUnauthenticatedGetByEmail(t *testing.T) {
	accountsRepoMock := new(mocks.AccountsRepo)
	accountsRepoMock.On("GetByEmail", testEmail).Return(nil, nil)
	cryptMock := crypt.New(crypt.NewConfig())
	authServiceClientMock := new(auth.MockAuthServiceClient)
	logIn := LogIn(accountsRepoMock, cryptMock, authServiceClientMock)

	ctx := context.Background()
	logInReq := &accounts.LogInRequest{
		AuthType: testAuthTypeEmail,
		Email:    testEmail,
		Password: testPassword,
	}
	logInRes, err := logIn(ctx, logInReq)
	assert.Nil(t, logInRes)
	assert.NotNil(t, err)
	assert.Equal(t, UnauthorizedError, err)
}

func TestLogInUnauthenticatedGetByID(t *testing.T) {
	accountsRepoMock := new(mocks.AccountsRepo)
	accountsRepoMock.On("GetByEmail", testEmail).Return(&model.Account{
		AccountID: testAccountID,
	}, nil)
	accountsRepoMock.On("GetByID", testAccountID).Return(nil, nil)
	cryptMock := crypt.NewMock()
	authServiceClientMock := new(auth.MockAuthServiceClient)
	logIn := LogIn(accountsRepoMock, cryptMock, authServiceClientMock)

	ctx := context.Background()
	logInReq := &accounts.LogInRequest{
		AuthType: testAuthTypeEmail,
		Email:    testEmail,
		Password: testPassword,
	}
	logInRes, err := logIn(ctx, logInReq)
	assert.Nil(t, logInRes)
	assert.NotNil(t, err)
	assert.Equal(t, UnauthorizedError, err)
}

func TestLogInUnauthorizedIncorrectPassword(t *testing.T) {
	accountsRepoMock := new(mocks.AccountsRepo)
	accountsRepoMock.On("GetByEmail", testEmail).Return(&model.Account{
		AccountID: testAccountID,
	}, nil)
	accountsRepoMock.On("GetByID", testAccountID).Return(&model.Account{}, nil)
	cryptMock := new(mocks.Crypt)
	cryptMock.On("CheckHashedPassword", "", testPassword).Return(false)
	authServiceClientMock := new(auth.MockAuthServiceClient)
	logIn := LogIn(accountsRepoMock, cryptMock, authServiceClientMock)

	ctx := context.Background()
	logInReq := &accounts.LogInRequest{
		AuthType: testAuthTypeEmail,
		Email:    testEmail,
		Password: testPassword,
	}
	logInRes, err := logIn(ctx, logInReq)
	assert.Nil(t, logInRes)
	assert.NotNil(t, err)
	assert.Equal(t, UnauthorizedError, err)
}

func TestLogInAuthError(t *testing.T) {
	accountsRepoMock := new(mocks.AccountsRepo)
	accountsRepoMock.On("GetByEmail", testEmail).Return(&model.Account{
		AccountID: testAccountID,
	}, nil)
	accountsRepoMock.On("GetByID", testAccountID).Return(&model.Account{
		AccountID:      testAccountID,
		HashedPassword: testHashedPassword,
	}, nil)
	cryptMock := new(mocks.Crypt)
	cryptMock.On("CheckHashedPassword", testHashedPassword, testPassword).Return(true)
	authServiceClientMock := new(auth.MockAuthServiceClient)
	ctx := context.Background()
	authServiceClientMock.On("Auth", ctx, &auth.AuthRequest{
		AccountId: testAccountID,
	}).Return(nil, jwt.ErrSignatureInvalid)
	logIn := LogIn(accountsRepoMock, cryptMock, authServiceClientMock)

	logInReq := &accounts.LogInRequest{
		AuthType: testAuthTypeEmail,
		Email:    testEmail,
		Password: testPassword,
	}
	logInRes, err := logIn(ctx, logInReq)
	assert.Nil(t, logInRes)
	assert.NotNil(t, err)
	assert.Equal(t, jwt.ErrSignatureInvalid, err)
}

func TestLogIn(t *testing.T) {
	accountsRepoMock := repo.NewMockAccountsRepo()
	cryptMock := crypt.NewMock()
	authServiceClientMock := new(auth.MockAuthServiceClient)
	authServiceClientMock.On("Auth", mock.Anything, mock.Anything).Return(&auth.AuthResponse{
		AccountId:    mock.Anything,
		AccessToken:  "test_access_token",
		RefreshToken: "test_refresh_token",
		ExpiresIn:    3600,
	}, nil)
	register := Register(accountsRepoMock, cryptMock)
	logIn := LogIn(accountsRepoMock, cryptMock, authServiceClientMock)

	ctx := context.Background()
	registerReq := &accounts.RegisterRequest{
		AuthType: testAuthTypeEmail,
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

	err = accountsRepoMock.DeleteByID(registeredAccountID)
	assert.Nil(t, err)
}
