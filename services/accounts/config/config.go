package config

import baseConfig "github.com/taeho-io/family/services/base/config"

type IFace interface {
	baseConfig.IFace

	Settings() Settings
}

type AccountsServiceConfig struct {
	baseConfig.IFace

	settings Settings
}

func New(settings Settings) (cfg IFace) {
	return &AccountsServiceConfig{
		IFace:    baseConfig.New(serviceName),
		settings: settings,
	}
}

func NewMock() (cfg IFace) {
	return &AccountsServiceConfig{
		IFace:    baseConfig.New(serviceName),
		settings: NewMockSettings(),
	}
}

func (c *AccountsServiceConfig) Settings() Settings {
	return c.settings
}
