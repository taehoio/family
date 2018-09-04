package repo

import "github.com/taeho-io/family/services/base"

type NotesRepoConfig interface {
	base.DynamodbRepoConfig
}

type defaultNotesRepoConfig struct {
	NotesRepoConfig

	fullTableName string
}

func NewNotesRepoConfig(fullTableName string) NotesRepoConfig {
	return &defaultNotesRepoConfig{
		fullTableName: fullTableName,
	}
}

func NewMockNotesRepoConfig() NotesRepoConfig {
	return NewNotesRepoConfig("family-development-notes-notes")
}

func (c *defaultNotesRepoConfig) FullTableName() string {
	return c.fullTableName
}
