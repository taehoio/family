package repo

import "github.com/taeho-io/family/services/base"

type TodosRepoConfig interface {
	base.DynamodbRepoConfig
}

type DefaultTogosRepoConfig struct {
	TodosRepoConfig

	fullTableName string
}

func NewTodosRepoConfig(fullTableName string) TodosRepoConfig {
	return &DefaultTogosRepoConfig{
		fullTableName: fullTableName,
	}
}

func NewMockTodosRepoConfig() TodosRepoConfig {
	return NewTodosRepoConfig("family-development-todos-todos")
}

func (c *DefaultTogosRepoConfig) FullTableName() string {
	return c.fullTableName
}
