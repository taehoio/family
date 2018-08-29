package base

type DynamodbRepoConfig interface {
	FullTableName() string
}
