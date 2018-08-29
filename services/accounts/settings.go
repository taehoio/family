package accounts

import "os"

const (
	serviceName                      = "accounts"
	defaultDynamodbAccountsTableName = "accounts"
)

type Settings struct {
	DynamodbAccountsTableName string
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

func MockSettings() Settings {
	return Settings{
		DynamodbAccountsTableName: defaultDynamodbAccountsTableName,
	}
}
