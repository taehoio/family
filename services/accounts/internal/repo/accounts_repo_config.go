package repo

import "github.com/taeho-io/family/services/base"

type AccountsRepoConfig interface {
	base.DynamodbRepoConfig
}

type DefaultAccountsRepoConfig struct {
	AccountsRepoConfig

	fullTableName string
}

func NewAccountsRepoConfig(fullTableName string) AccountsRepoConfig {
	return &DefaultAccountsRepoConfig{
		fullTableName: fullTableName,
	}
}

func NewMockAccountRepoConfig() AccountsRepoConfig {
	return NewAccountsRepoConfig("family-development-accounts-accounts")
}

func (c *DefaultAccountsRepoConfig) FullTableName() string {
	return c.fullTableName
}
