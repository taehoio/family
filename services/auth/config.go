package auth

import "github.com/taeho-io/family/services/base"

type Config interface {
	base.Config

	Settings() Settings
}

type DefaultConfig struct {
	base.Config

	settings Settings
}

func NewConfig(settings Settings) (cfg Config) {
	return &DefaultConfig{
		Config:   base.NewConfig(srvName),
		settings: settings,
	}
}

func NewMockConfig() (cfg Config) {
	return NewConfig(MockSettings())
}

func (c *DefaultConfig) Settings() Settings {
	return c.settings
}
