package handlers

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	mockAuth "github.com/taeho-io/family/idl/generated/go/pb/family/mocks/auth"
	"github.com/taeho-io/family/services/accounts/crypt"
	mockCrypt "github.com/taeho-io/family/services/accounts/mocks/crypt"
	mockAccountsRepo "github.com/taeho-io/family/services/accounts/mocks/repos/accounts_repo"
	"github.com/taeho-io/family/services/accounts/models"
	"github.com/taeho-io/family/services/accounts/repos/accounts_repo"
)

func TestLogInUnauthenticatedGetByEmail(t *testing.T) {
	accountsTableMock := new(mockAccountsRepo.IFace)
	accountsTableMock.On("GetByEmail", testEmail).Return(nil, nil)
	cryptMock := crypt.New()
	authServiceClientMock := new(mockAuth.AuthServiceClient)
	logIn := LogIn(accountsTableMock, cryptMock, authServiceClientMock)

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
	accountsTableMock := new(mockAccountsRepo.IFace)
	accountsTableMock.On("GetByEmail", testEmail).Return(&models.Account{
		AccountID: testAccountID,
	}, nil)
	accountsTableMock.On("GetByID", testAccountID).Return(nil, nil)
	cryptMock := crypt.New()
	authServiceClientMock := new(mockAuth.AuthServiceClient)
	logIn := LogIn(accountsTableMock, cryptMock, authServiceClientMock)

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
	accountsTableMock := new(mockAccountsRepo.IFace)
	accountsTableMock.On("GetByEmail", testEmail).Return(&models.Account{
		AccountID: testAccountID,
	}, nil)
	accountsTableMock.On("GetByID", testAccountID).Return(&models.Account{}, nil)
	cryptMock := new(mockCrypt.IFace)
	cryptMock.On("CheckHashedPassword", "", testPassword).Return(false)
	authServiceClientMock := new(mockAuth.AuthServiceClient)
	logIn := LogIn(accountsTableMock, cryptMock, authServiceClientMock)

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
	accountsTableMock := new(mockAccountsRepo.IFace)
	accountsTableMock.On("GetByEmail", testEmail).Return(&models.Account{
		AccountID: testAccountID,
	}, nil)
	accountsTableMock.On("GetByID", testAccountID).Return(&models.Account{
		AccountID:      testAccountID,
		HashedPassword: testHashedPassword,
	}, nil)
	cryptMock := new(mockCrypt.IFace)
	cryptMock.On("CheckHashedPassword", testHashedPassword, testPassword).Return(true)
	authServiceClientMock := new(mockAuth.AuthServiceClient)
	ctx := context.Background()
	authServiceClientMock.On("Auth", ctx, &auth.AuthRequest{
		AccountId: testAccountID,
	}).Return(nil, jwt.ErrSignatureInvalid)
	logIn := LogIn(accountsTableMock, cryptMock, authServiceClientMock)

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
	accountsTableMock := accounts_repo.NewMock()
	cryptMock := crypt.New()
	authServiceClientMock := new(mockAuth.AuthServiceClient)
	authServiceClientMock.On("Auth", mock.Anything, mock.Anything).Return(&auth.AuthResponse{
		AccountId:    mock.Anything,
		AccessToken:  "test_access_token",
		RefreshToken: "test_refresh_token",
		ExpiresIn:    3600,
	}, nil)
	register := Register(accountsTableMock, cryptMock)
	logIn := LogIn(accountsTableMock, cryptMock, authServiceClientMock)

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

	err = accountsTableMock.DeleteByID(registeredAccountID)
	assert.Nil(t, err)
}
