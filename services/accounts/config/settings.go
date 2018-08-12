package config

import "os"

const (
	serviceName                      = "accounts"
	defaultDynamodbAccountsTableName = "accounts"
)

type Settings struct {
	DynamodbAccountsTableName      string
	DynamodbAccountsEmailTableName string
}

func NewSettings() Settings {
	dynamodbAccountsTableName := os.Getenv("DynamoDbAccountsTableName")
	if dynamodbAccountsTableName == "" {
		dynamodbAccountsTableName = defaultDynamodbAccountsTableName
	}

	return Settings{
		DynamodbAccountsTableName: dynamodbAccountsTableName,
	}
}

func NewMockSettings() Settings {
	return Settings{
		DynamodbAccountsTableName: defaultDynamodbAccountsTableName,
	}
}
