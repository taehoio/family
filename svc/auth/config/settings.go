package config

import (
	"os"
	"time"
)

const (
	namespace = "auth"
)

type Settings struct {
	TokenIssuer                  string
	AccessTokenExpiringDuration  time.Duration
	RefreshTokenExpiringDuration time.Duration
}

func NewSettings() Settings {
	return Settings{
		TokenIssuer:                  os.Getenv("TOKEN_ISSUER"),
		AccessTokenExpiringDuration:  time.Hour,
		RefreshTokenExpiringDuration: time.Hour * 24 * 365,
	}
}

func NewMockSettings() Settings {
	return Settings{
		TokenIssuer:                  "MOCK_TOKEN_ISSUER",
		AccessTokenExpiringDuration:  time.Hour,
		RefreshTokenExpiringDuration: time.Hour * 24 * 365,
	}
}
