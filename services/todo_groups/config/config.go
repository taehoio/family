package config

import baseConfig "github.com/taeho-io/family/services/base/config"

type IFace interface {
	baseConfig.IFace

	Settings() Settings
}

type TodoGroupsServiceConfig struct {
	baseConfig.IFace

	settings Settings
}

func New(settings Settings) (cfg IFace) {
	return &TodoGroupsServiceConfig{
		IFace:    baseConfig.New(serviceName),
		settings: settings,
	}

}

func NewMock() (cfg IFace) {
	return &TodoGroupsServiceConfig{
		IFace:    baseConfig.New(serviceName),
		settings: NewMockSettings(),
	}
}

func (t *TodoGroupsServiceConfig) Settings() Settings {
	return t.settings
}
