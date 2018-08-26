package token

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/taeho-io/family/services/auth/config"
	"github.com/taeho-io/family/services/auth/jwt"
)

const (
	testAccountId = "test_account_id"
)

func TestNewAccessToken(t *testing.T) {
	cfg := config.New(config.NewMockSettings())
	tokenSvc := New(cfg)
	token, err := tokenSvc.NewAccessToken(testAccountId)
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestNewRefreshToken(t *testing.T) {
	cfg := config.New(config.NewMockSettings())
	tokenSvc := New(cfg)
	token, err := tokenSvc.NewRefreshToken(testAccountId)
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestValidateToken(t *testing.T) {
	cfg := config.New(config.NewMockSettings())
	tokenSvc := New(cfg)
	token, _ := tokenSvc.NewAccessToken(testAccountId)
	err := tokenSvc.ValidateToken(token)
	assert.Nil(t, err)
}

func TestToClaims(t *testing.T) {
	cfg := config.New(config.NewMockSettings())
	tokenSvc := New(cfg)
	claims := tokenSvc.toClaims(&jwt.Claims{
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
	tokenSvc := New(cfg)
	token, _ := tokenSvc.NewAccessToken(testAccountId)
	claims, err := tokenSvc.ParseToken(token)
	assert.Nil(t, err)
	assert.NotNil(t, claims)
}

func TestParseToken_Error(t *testing.T) {
	cfg := config.New(config.NewMockSettings())
	tokenSvc := New(cfg)
	token := "invalid_token"
	claims, err := tokenSvc.ParseToken(token)
	assert.NotNil(t, err)
	assert.Nil(t, claims)
}
