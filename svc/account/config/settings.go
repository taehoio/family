package config

import "os"

const (
	srvName = "account"
)

type Settings struct {
	DynamodbAccountTableName      string
	DynamodbAccountEmailTableName string
}

func NewSettings() Settings {
	return Settings{
		DynamodbAccountTableName:      os.Getenv("DynamoDbAccountTableName"),
		DynamodbAccountEmailTableName: os.Getenv("DynamoDbAccountEmailTableName"),
	}
}

func NewMockSettings() Settings {
	return Settings{
		DynamodbAccountTableName:      "account",
		DynamodbAccountEmailTableName: "account_email",
	}
}
