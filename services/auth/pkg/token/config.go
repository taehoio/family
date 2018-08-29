package token

import "time"

type Config interface {
	TokenIssuer() string
	AccessTokenExpiringDuration() time.Duration
	RefreshTokenExpiringDuration() time.Duration
}

type DefaultConfig struct {
	Config

	tokenIssuer                  string
	accessTokenExpiringDuration  time.Duration
	refreshTokenExpiringDuration time.Duration
}

func NewConfig(tokenIssuer string, accessTokenExpiringDuration, refreshTokenExpiringDuration time.Duration) Config {
	return &DefaultConfig{
		tokenIssuer:                  tokenIssuer,
		accessTokenExpiringDuration:  accessTokenExpiringDuration,
		refreshTokenExpiringDuration: refreshTokenExpiringDuration,
	}
}

func MockConfig() Config {
	return NewConfig("MOCK_TOKEN_ISSUER", time.Hour, time.Hour*24*365)
}

func (c *DefaultConfig) TokenIssuer() string {
	return c.tokenIssuer
}

func (c *DefaultConfig) AccessTokenExpiringDuration() time.Duration {
	return c.accessTokenExpiringDuration
}

func (c *DefaultConfig) RefreshTokenExpiringDuration() time.Duration {
	return c.refreshTokenExpiringDuration
}
