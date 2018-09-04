package repo

import "github.com/taeho-io/family/services/base"

type TodosRepoConfig interface {
	base.DynamodbRepoConfig
}

type defaultTodosRepoConfig struct {
	TodosRepoConfig

	fullTableName string
}

func NewTodosRepoConfig(fullTableName string) TodosRepoConfig {
	return &defaultTodosRepoConfig{
		fullTableName: fullTableName,
	}
}

func NewMockTodosRepoConfig() TodosRepoConfig {
	return NewTodosRepoConfig("family-development-todos-todos")
}

func (c *defaultTodosRepoConfig) FullTableName() string {
	return c.fullTableName
}
