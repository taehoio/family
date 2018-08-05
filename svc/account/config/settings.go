package config

import "os"

const (
	srvName = "account"
)

type Settings struct {
	AwsDynamodbRegion             string
	DynamodbAccountTableName      string
	DynamodbAccountEmailTableName string
}

func NewSettings() Settings {
	return Settings{
		AwsDynamodbRegion:             os.Getenv("AWS_DEFAULT_REGION"),
		DynamodbAccountTableName:      os.Getenv("DynamoDbAccountTableName"),
		DynamodbAccountEmailTableName: os.Getenv("DynamoDbAccountEmailTableName"),
	}
}

func NewMockSettings() Settings {
	return Settings{
		AwsDynamodbRegion:             "us-west-2",
		DynamodbAccountTableName:      "account",
		DynamodbAccountEmailTableName: "account_email",
	}
}
