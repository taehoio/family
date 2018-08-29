package repo

import "github.com/taeho-io/family/services/base"

type TodoGroupsRepoConfig interface {
	base.DynamodbRepoConfig
}

type DefaultTogoGroupsRepoConfig struct {
	TodoGroupsRepoConfig

	fullTableName string
}

func NewTodoGroupsRepoConfig(fullTableName string) TodoGroupsRepoConfig {
	return &DefaultTogoGroupsRepoConfig{
		fullTableName: fullTableName,
	}
}

func NewMockTodoGroupsRepoConfig() TodoGroupsRepoConfig {
	return NewTodoGroupsRepoConfig("family-development-todogroups-todo_groups")
}

func (c *DefaultTogoGroupsRepoConfig) FullTableName() string {
	return c.fullTableName
}
