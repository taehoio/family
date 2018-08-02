package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"errors"

	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/svc/auth/config"
	"github.com/taeho-io/family/svc/auth/mocks"
	"github.com/taeho-io/family/svc/auth/token"
)

func TestRefreshHandler(t *testing.T) {
	ctx := context.Background()
	settings := config.NewSettings()
	cfg := config.New(settings)
	tokenSrv := token.NewJwtToken(cfg)
	refreshToken, _ := tokenSrv.NewRefreshToken(testAccountId)
	req := &auth.RefreshRequest{
		RefreshToken: refreshToken,
	}
	res, err := Refresh(cfg, tokenSrv)(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestRefreshHandler_Error_InvalidRefreshToken(t *testing.T) {
	ctx := context.Background()
	settings := config.NewSettings()
	cfg := config.New(settings)
	tokenSrv := token.NewJwtToken(cfg)
	req := &auth.RefreshRequest{
		RefreshToken: "invalid_token",
	}
	res, err := Refresh(cfg, tokenSrv)(ctx, req)
	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestRefreshHandler_NewAccessToken_Error(t *testing.T) {
	ctx := context.Background()
	settings := config.NewSettings()
	cfg := config.New(settings)
	refreshToken, _ := token.NewJwtToken(cfg).NewRefreshToken(testAccountId)
	req := &auth.RefreshRequest{
		RefreshToken: refreshToken,
	}
	tokenSrv := new(mocks.Token)
	tokenSrv.On("NewAccessToken", testAccountId).Return("", errors.New("failed"))
	_, err := Refresh(cfg, tokenSrv)(ctx, req)
	assert.NotNil(t, err)
}
