package config

import "github.com/taeho-io/family/services/base/config"

type IFace interface {
	config.IFace

	Settings() Settings
}

type TodoGroupsServiceConfig struct {
	config.IFace

	settings Settings
}

func New(settings Settings) (cfg IFace) {
	return &TodoGroupsServiceConfig{
		IFace:    config.New(serviceName),
		settings: settings,
	}

}

func NewMock() (cfg IFace) {
	return &TodoGroupsServiceConfig{
		IFace:    config.New(serviceName),
		settings: NewMockSettings(),
	}
}

func (t *TodoGroupsServiceConfig) Settings() Settings {
	return t.settings
}
