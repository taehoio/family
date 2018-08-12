package config

import baseConfig "github.com/taeho-io/family/services/base/config"

type IFace interface {
	baseConfig.IFace

	Settings() Settings
}

type TodosServiceConfig struct {
	baseConfig.IFace

	settings Settings
}

func New(settings Settings) (cfg IFace) {
	return &TodosServiceConfig{
		IFace:    baseConfig.New(serviceName),
		settings: settings,
	}

}

func NewMock() (cfg IFace) {
	return &TodosServiceConfig{
		IFace:    baseConfig.New(serviceName),
		settings: NewMockSettings(),
	}
}

func (t *TodosServiceConfig) Settings() Settings {
	return t.settings
}
