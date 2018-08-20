package config

import "github.com/taeho-io/family/services/base/config"

type IFace interface {
	config.IFace

	Settings() Settings
}

type AccountsServiceConfig struct {
	config.IFace

	settings Settings
}

func New(settings Settings) (cfg IFace) {
	return &AccountsServiceConfig{
		IFace:    config.New(serviceName),
		settings: settings,
	}
}

func NewMock() (cfg IFace) {
	return &AccountsServiceConfig{
		IFace:    config.New(serviceName),
		settings: NewMockSettings(),
	}
}

func (c *AccountsServiceConfig) Settings() Settings {
	return c.settings
}
