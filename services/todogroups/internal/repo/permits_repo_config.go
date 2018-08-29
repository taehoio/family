package repo

import "github.com/taeho-io/family/services/base"

type TodoGroupPermitsRepoConfig interface {
	base.DynamodbRepoConfig
}

type DefaultTodoGroupPermitsRepoConfig struct {
	TodoGroupPermitsRepoConfig

	fullTableName string
}

func NewTodoGroupPermitsRepoConfig(fullTableName string) TodoGroupPermitsRepoConfig {
	return &DefaultTodoGroupPermitsRepoConfig{
		fullTableName: fullTableName,
	}
}

func NewMockPermitsRepoConfig() TodoGroupPermitsRepoConfig {
	return NewTodoGroupPermitsRepoConfig("family-development-todogroups-todo_group_permits")
}

func (c *DefaultTodoGroupPermitsRepoConfig) FullTableName() string {
	return c.fullTableName
}
