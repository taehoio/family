package todos

import (
	"github.com/taeho-io/family/services/base"
)

type Config interface {
	base.Config

	Settings() Settings
}

type defaultConfig struct {
	base.Config

	settings Settings
}

func NewConfig(settings Settings) Config {
	return &defaultConfig{
		Config:   base.NewConfig(serviceName),
		settings: settings,
	}
}

func NewMockConfig() Config {
	return &defaultConfig{}
}

func (t *defaultConfig) Settings() Settings {
	return t.settings
}
