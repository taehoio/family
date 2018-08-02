package token

import (
	"time"

	"github.com/taeho-io/family/svc/auth/config"
	"github.com/taeho-io/family/svc/auth/jwt"
)

type Token interface {
	NewAccessToken(string) (string, error)
	NewRefreshToken(string) (string, error)
	ValidateToken(string) error
	ParseToken(string) (*Claims, error)
}

type Claims struct {
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`
}

type JwtToken struct {
	config config.IFace

	Token
}

func NewJwtToken(cfg config.IFace) *JwtToken {
	return &JwtToken{
		config: cfg,
	}
}

func (t *JwtToken) NewAccessToken(accountId string) (string, error) {
	settings := t.config.Settings()
	claimsForAccessToken := &jwt.Claims{
		Issuer:    settings.TokenIssuer,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(settings.AccessTokenExpiringDuration).Unix(),
		Audience:  accountId,
		Subject:   "AuthToken",
	}
	return jwt.NewWithClaims(claimsForAccessToken)
}

func (t *JwtToken) NewRefreshToken(accountId string) (string, error) {
	settings := t.config.Settings()
	claimsForRefreshToken := &jwt.Claims{
		Issuer:    settings.TokenIssuer,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(settings.AccessTokenExpiringDuration).Unix(),
		Audience:  accountId,
		Subject:   "RefreshToken",
	}
	return jwt.NewWithClaims(claimsForRefreshToken)
}

func (t *JwtToken) ValidateToken(tokenString string) error {
	return jwt.ValidateToken(tokenString)
}

func (t *JwtToken) toClaims(claims *jwt.Claims) *Claims {
	return &Claims{
		Audience:  claims.Audience,
		ExpiresAt: claims.ExpiresAt,
		Id:        claims.Id,
		IssuedAt:  claims.IssuedAt,
		Issuer:    claims.Issuer,
		NotBefore: claims.NotBefore,
		Subject:   claims.Subject,
	}
}

func (t *JwtToken) ParseToken(tokenString string) (*Claims, error) {
	jwtClaims, err := jwt.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	return t.toClaims(jwtClaims), nil
}
