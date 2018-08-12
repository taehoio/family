package config

import srvConfig "github.com/taeho-io/family/services/base/config"

type IFace interface {
	srvConfig.IFace

	Settings() Settings
}

type AuthSvcConfig struct {
	srvConfig.IFace

	settings Settings
}

func New(settings Settings) (cfg IFace) {
	return &AuthSvcConfig{
		IFace:    srvConfig.New(srvName),
		settings: settings,
	}
}

func NewMock() (cfg IFace) {
	return New(NewMockSettings())
}

func (c *AuthSvcConfig) Settings() Settings {
	return c.settings
}
