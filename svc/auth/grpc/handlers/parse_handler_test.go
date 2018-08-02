package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/svc/auth/config"
	"github.com/taeho-io/family/svc/auth/token"
)

func TestParseHandler(t *testing.T) {
	ctx := context.Background()
	settings := config.NewSettings()
	cfg := config.New(settings)
	tokenSrv := token.NewJwtToken(cfg)
	accessToken, _ := tokenSrv.NewAccessToken(testAccountId)
	req := &auth.ParseRequest{
		AccessToken: accessToken,
	}
	res, err := Parse(tokenSrv)(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestParseHandler_Error(t *testing.T) {
	ctx := context.Background()
	settings := config.NewSettings()
	cfg := config.New(settings)
	tokenSrv := token.NewJwtToken(cfg)
	req := &auth.ParseRequest{
		AccessToken: "invalid_token",
	}
	res, err := Parse(tokenSrv)(ctx, req)
	assert.NotNil(t, err)
	assert.Nil(t, res)
}
