package config

import srvConfig "github.com/taeho-io/family/svc/srv/config"

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
		IFace:    srvConfig.New(namespace),
		settings: settings,
	}
}

func (c *AuthSvcConfig) Settings() Settings {
	return c.settings
}
