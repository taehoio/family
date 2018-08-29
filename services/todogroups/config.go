package todogroups

import (
	"github.com/taeho-io/family/services/base"
)

type Config interface {
	base.Config

	Settings() Settings
}

type TodoGroupsServiceConfig struct {
	base.Config

	settings Settings
}

func NewConfig(settings Settings) (cfg Config) {
	return &TodoGroupsServiceConfig{
		Config:   base.NewConfig(serviceName),
		settings: settings,
	}

}

func NewMockConfig() (cfg Config) {
	return &TodoGroupsServiceConfig{
		Config:   base.NewConfig(serviceName),
		settings: NewMockSettings(),
	}
}

func (t *TodoGroupsServiceConfig) Settings() Settings {
	return t.settings
}
