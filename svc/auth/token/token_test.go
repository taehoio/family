package token

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/taeho-io/family/svc/auth/config"
	"github.com/taeho-io/family/svc/auth/jwt"
)

const (
	testAccountId = "test_account_id"
)

func TestNewAccessToken(t *testing.T) {
	cfg := config.New(config.NewMockSettings())
	tokenSrv := New(cfg)
	token, err := tokenSrv.NewAccessToken(testAccountId)
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestNewRefreshToken(t *testing.T) {
	cfg := config.New(config.NewMockSettings())
	tokenSrv := New(cfg)
	token, err := tokenSrv.NewRefreshToken(testAccountId)
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestValidateToken(t *testing.T) {
	cfg := config.New(config.NewMockSettings())
	tokenSrv := New(cfg)
	token, _ := tokenSrv.NewAccessToken(testAccountId)
	err := tokenSrv.ValidateToken(token)
	assert.Nil(t, err)
}

func TestToClaims(t *testing.T) {
	cfg := config.New(config.NewMockSettings())
	tokenSrv := New(cfg)
	claims := tokenSrv.toClaims(&jwt.Claims{
		Audience:  "",
		ExpiresAt: 0,
		Id:        "",
		IssuedAt:  0,
		Issuer:    "",
		NotBefore: 0,
		Subject:   "",
	})

	expected := &Claims{
		Audience:  "",
		ExpiresAt: 0,
		Id:        "",
		IssuedAt:  0,
		Issuer:    "",
		NotBefore: 0,
		Subject:   "",
	}
	assert.EqualValues(t, expected, claims)
}

func TestParseToken(t *testing.T) {
	cfg := config.New(config.NewMockSettings())
	tokenSrv := New(cfg)
	token, _ := tokenSrv.NewAccessToken(testAccountId)
	claims, err := tokenSrv.ParseToken(token)
	assert.Nil(t, err)
	assert.NotNil(t, claims)
}

func TestParseToken_Error(t *testing.T) {
	cfg := config.New(config.NewMockSettings())
	tokenSrv := New(cfg)
	token := "invalid_token"
	claims, err := tokenSrv.ParseToken(token)
	assert.NotNil(t, err)
	assert.Nil(t, claims)
}
