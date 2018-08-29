package token

import (
	"time"

	"github.com/taeho-io/family/services/auth/pkg/jwt"
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
	Token

	cfg Config
}

func New(cfg Config) Token {
	return &JwtToken{
		cfg: cfg,
	}
}

func Mock() Token {
	return New(MockConfig())
}

func (t *JwtToken) NewAccessToken(accountId string) (string, error) {
	claimsForAccessToken := &jwt.Claims{
		Issuer:    t.cfg.TokenIssuer(),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(t.cfg.AccessTokenExpiringDuration()).Unix(),
		Audience:  accountId,
		Subject:   "AuthToken",
	}
	return jwt.NewWithClaims(claimsForAccessToken)
}

func (t *JwtToken) NewRefreshToken(accountId string) (string, error) {
	claimsForRefreshToken := &jwt.Claims{
		Issuer:    t.cfg.TokenIssuer(),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(t.cfg.RefreshTokenExpiringDuration()).Unix(),
		Audience:  accountId,
		Subject:   "RefreshToken",
	}
	return jwt.NewWithClaims(claimsForRefreshToken)
}

func (t *JwtToken) ValidateToken(tokenString string) error {
	return jwt.ValidateToken(tokenString)
}

func (t *JwtToken) ParseToken(tokenString string) (*Claims, error) {
	jwtClaims, err := jwt.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	return toClaims(jwtClaims), nil
}

func toClaims(claims *jwt.Claims) *Claims {
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
