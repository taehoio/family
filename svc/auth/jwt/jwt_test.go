package jwt

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

var (
	claims = &Claims{
		Id:        "test_account_id",
		Issuer:    "test_issuer",
		Audience:  "test_audience",
		ExpiresAt: time.Now().Unix(),
	}
	validTokenString = ""
)

func TestNewWithClaims(t *testing.T) {
	tokenString, err := NewWithClaims(claims)
	assert.NotNil(t, tokenString)
	assert.Nil(t, err)

	validTokenString = tokenString
}

func TestKeyFunc_WrongSigningMethod_Error(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodNone, jwt.StandardClaims(*claims))
	_, err := keyFunc(token)
	assert.NotNil(t, err)
}

func TestValidateToken(t *testing.T) {
	err := ValidateToken("invalid_token")
	assert.NotNil(t, err)
}

func TestParseToken(t *testing.T) {
	claims, err := ParseToken(validTokenString)
	assert.NotNil(t, claims)
	assert.Nil(t, err)
}

func TestParseToken_Error(t *testing.T) {
	claims, err := ParseToken("invalid_token")
	assert.Nil(t, claims)
	assert.NotNil(t, err)
}
