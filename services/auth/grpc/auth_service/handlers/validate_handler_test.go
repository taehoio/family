package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/services/auth/config"
	"github.com/taeho-io/family/services/auth/token"
)

func TestValidateHandler(t *testing.T) {
	ctx := context.Background()
	settings := config.NewSettings()
	cfg := config.New(settings)
	tokenSvc := token.New(cfg)
	accessToken, _ := tokenSvc.NewAccessToken(testAccountId)
	req := &auth.ValidateRequest{
		AccessToken: accessToken,
	}
	res, err := Validate(tokenSvc)(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
