package config

import baseConfig "github.com/taeho-io/family/services/base/config"

type IFace interface {
	baseConfig.IFace

	Settings() Settings
}

type AuthServiceConfig struct {
	baseConfig.IFace

	settings Settings
}

func New(settings Settings) (cfg IFace) {
	return &AuthServiceConfig{
		IFace:    baseConfig.New(srvName),
		settings: settings,
	}
}

func NewMock() (cfg IFace) {
	return New(NewMockSettings())
}

func (c *AuthServiceConfig) Settings() Settings {
	return c.settings
}
