package config

import srvConfig "github.com/taeho-io/family/svc/srv/config"

type IFace interface {
	srvConfig.IFace

	Settings() Settings
}

type AccountSvcConfig struct {
	srvConfig.IFace

	settings Settings
}

func New(settings Settings) (cfg IFace) {
	return &AccountSvcConfig{
		IFace:    srvConfig.New(srvName),
		settings: settings,
	}
}

func NewMock() (cfg IFace) {
	return &AccountSvcConfig{
		IFace:    srvConfig.New(srvName),
		settings: NewMockSettings(),
	}
}

func (c *AccountSvcConfig) Settings() Settings {
	return c.settings
}
