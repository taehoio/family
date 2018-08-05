package handlers

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/svc/auth/config"
	"github.com/taeho-io/family/svc/auth/mocks"
	"github.com/taeho-io/family/svc/auth/token"
)

const (
	testAccountId = "test_account_id"
)

func TestAuthHandler(t *testing.T) {
	ctx := context.Background()
	req := &auth.AuthRequest{
		AccountId: testAccountId,
	}
	settings := config.NewSettings()
	cfg := config.New(settings)
	tokenSrv := token.New(cfg)
	res, err := Auth(cfg, tokenSrv)(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	accessToken, _ := tokenSrv.NewAccessToken(testAccountId)
	refreshToken, _ := tokenSrv.NewRefreshToken(testAccountId)
	expected := &auth.AuthResponse{
		AccountId:    testAccountId,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(settings.AccessTokenExpiringDuration.Seconds()),
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
	settings := config.NewSettings()
	cfg := config.New(settings)
	tokenSrv := new(mocks.Token)
	tokenSrv.On("NewAccessToken", testAccountId).Return("", errors.New("failed"))
	_, err := Auth(cfg, tokenSrv)(ctx, req)
	assert.NotNil(t, err)
}

func TestAuthHandler_NewRefreshToken_Error(t *testing.T) {
	ctx := context.Background()
	req := &auth.AuthRequest{
		AccountId: testAccountId,
	}
	settings := config.NewSettings()
	cfg := config.New(settings)
	tokenSrv := new(mocks.Token)
	tokenSrv.On("NewAccessToken", testAccountId).Return("token", nil)
	tokenSrv.On("NewRefreshToken", testAccountId).Return("", errors.New("failed"))
	_, err := Auth(cfg, tokenSrv)(ctx, req)
	assert.NotNil(t, err)
}
