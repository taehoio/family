package config

import baseConfig "github.com/taeho-io/family/services/base/config"

type IFace interface {
	baseConfig.Config

	Settings() Settings
}

type TodosServiceConfig struct {
	baseConfig.Config

	settings Settings
}

func New(settings Settings) (cfg IFace) {
	return &TodosServiceConfig{
		Config:   baseConfig.NewConfig(serviceName),
		settings: settings,
	}

}

func NewMock() (cfg IFace) {
	return &TodosServiceConfig{
		Config:   baseConfig.NewConfig(serviceName),
		settings: NewMockSettings(),
	}
}

func (t *TodosServiceConfig) Settings() Settings {
	return t.settings
}
