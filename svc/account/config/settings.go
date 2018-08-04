package config

import "os"

const (
	srvName = "account"
)

type Settings struct {
	AwsDynamodbRegion        string
	DynamodbAccountTableName string
}

func NewSettings() Settings {
	return Settings{
		AwsDynamodbRegion:        os.Getenv("AWS_DEFAULT_REGION"),
		DynamodbAccountTableName: os.Getenv("DynamoDbAccountTableName"),
	}
}

func NewMockSettings() Settings {
	return Settings{
		AwsDynamodbRegion:        "us-west-2",
		DynamodbAccountTableName: "account",
	}
}
