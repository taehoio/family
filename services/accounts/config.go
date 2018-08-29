package accounts

import (
	"github.com/taeho-io/family/services/base"
)

type Config interface {
	base.Config

	Settings() Settings
}

type DefaultConfig struct {
	base.Config

	settings Settings
}

func NewConfig(settings Settings) Config {
	return &DefaultConfig{
		Config:   base.NewConfig(serviceName),
		settings: settings,
	}
}

func NewMockConfig() (cfg Config) {
	return &DefaultConfig{
		Config:   base.NewConfig(serviceName),
		settings: MockSettings(),
	}
}

func (c *DefaultConfig) Settings() Settings {
	return c.settings
}
