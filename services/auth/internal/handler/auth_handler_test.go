package handler

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/services/auth/mocks"
	"github.com/taeho-io/family/services/auth/pkg/token"
)

const (
	testAccountId                   = "test_account_id"
	testAccessTokenExpiringDuration = time.Hour
)

func TestAuthHandler(t *testing.T) {
	ctx := context.Background()
	req := &auth.AuthRequest{
		AccountId: testAccountId,
	}

	tokenSvc := token.Mock()

	res, err := Auth(testAccessTokenExpiringDuration, tokenSvc)(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	accessToken, _ := tokenSvc.NewAccessToken(testAccountId)
	refreshToken, _ := tokenSvc.NewRefreshToken(testAccountId)
	expected := &auth.AuthResponse{
		AccountId:    testAccountId,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(testAccessTokenExpiringDuration.Seconds()),
	}
	assert.Equal(t, res.AccountId, expected.AccountId)
	assert.Equal(t, res.AccessToken, expected.AccessToken)
	assert.Equal(t, res.RefreshToken, expected.RefreshToken)
	assert.Equal(t, res.ExpiresIn, expected.ExpiresIn)
}

func TestAuthHandler_NewAccessToken_Error(t *testing.T) {
	ctx := context.Background()
	req := &auth.AuthRequest{
		AccountId: testAccountId,
	}

	tokenSvc := new(mocks.Token)
	tokenSvc.On("NewAccessToken", testAccountId).Return("", errors.New("failed"))
	_, err := Auth(testAccessTokenExpiringDuration, tokenSvc)(ctx, req)
	assert.NotNil(t, err)
}

func TestAuthHandler_NewRefreshToken_Error(t *testing.T) {
	ctx := context.Background()
	req := &auth.AuthRequest{
		AccountId: testAccountId,
	}

	tokenSvc := new(mocks.Token)
	tokenSvc.On("NewAccessToken", testAccountId).Return("token", nil)
	tokenSvc.On("NewRefreshToken", testAccountId).Return("", errors.New("failed"))
	_, err := Auth(testAccessTokenExpiringDuration, tokenSvc)(ctx, req)
	assert.NotNil(t, err)
}
