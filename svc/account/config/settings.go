package config

import "os"

const (
	srvName                              = "account"
	defaultDynamodbAccountTableName      = "account"
	defaultDynamodbAccountEmailTableName = "account_email"
)

type Settings struct {
	DynamodbAccountTableName      string
	DynamodbAccountEmailTableName string
}

func NewSettings() Settings {
	dynamodbAccountTableName := os.Getenv("DynamoDbAccountTableName")
	if dynamodbAccountTableName == "" {
		dynamodbAccountTableName = defaultDynamodbAccountTableName
	}
	dynamodbAccountEmailTableName := os.Getenv("DynamoDbAccountEmailTableName")
	if dynamodbAccountEmailTableName == "" {
		dynamodbAccountEmailTableName = defaultDynamodbAccountEmailTableName
	}

	return Settings{
		DynamodbAccountTableName:      dynamodbAccountTableName,
		DynamodbAccountEmailTableName: dynamodbAccountEmailTableName,
	}
}

func NewMockSettings() Settings {
	return Settings{
		DynamodbAccountTableName:      defaultDynamodbAccountTableName,
		DynamodbAccountEmailTableName: defaultDynamodbAccountEmailTableName,
	}
}
