package config

import "os"

const (
	srvName                                = "accounts"
	defaultDynamodbAccountsTableName       = "accounts"
	defaultDynamodbAccountsEmailsTableName = "accounts_emails"
)

type Settings struct {
	DynamodbAccountTableName      string
	DynamodbAccountEmailTableName string
}

func NewSettings() Settings {
	dynamodbAccountsTableName := os.Getenv("DynamoDbAccountsTableName")
	if dynamodbAccountsTableName == "" {
		dynamodbAccountsTableName = defaultDynamodbAccountsTableName
	}
	dynamodbAccountsEmailsTableName := os.Getenv("DynamoDbAccountsEmailsTableName")
	if dynamodbAccountsEmailsTableName == "" {
		dynamodbAccountsEmailsTableName = defaultDynamodbAccountsEmailsTableName
	}

	return Settings{
		DynamodbAccountTableName:      dynamodbAccountsTableName,
		DynamodbAccountEmailTableName: dynamodbAccountsEmailsTableName,
	}
}

func NewMockSettings() Settings {
	return Settings{
		DynamodbAccountTableName:      defaultDynamodbAccountsTableName,
		DynamodbAccountEmailTableName: defaultDynamodbAccountsEmailsTableName,
	}
}
