package handler

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/services/auth/mocks"
	"github.com/taeho-io/family/services/auth/pkg/token"
)

func TestRefreshHandler(t *testing.T) {
	ctx := context.Background()

	tokenSvc := token.Mock()

	refreshToken, _ := tokenSvc.NewRefreshToken(testAccountId)
	req := &auth.RefreshRequest{
		RefreshToken: refreshToken,
	}
	res, err := Refresh(testAccessTokenExpiringDuration, tokenSvc)(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestRefreshHandler_Error_InvalidRefreshToken(t *testing.T) {
	ctx := context.Background()

	tokenSvc := token.Mock()

	req := &auth.RefreshRequest{
		RefreshToken: "invalid_token",
	}
	res, err := Refresh(testAccessTokenExpiringDuration, tokenSvc)(ctx, req)
	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestRefreshHandler_NewAccessToken_Error(t *testing.T) {
	ctx := context.Background()

	refreshToken, _ := token.Mock().NewRefreshToken(testAccountId)
	req := &auth.RefreshRequest{
		RefreshToken: refreshToken,
	}
	tokenSvc := new(mocks.Token)
	tokenSvc.On("NewAccessToken", testAccountId).Return("", errors.New("failed"))
	_, err := Refresh(testAccessTokenExpiringDuration, tokenSvc)(ctx, req)
	assert.NotNil(t, err)
}
