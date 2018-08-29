package todos

import (
	"os"
)

const (
	serviceName                   = "todos"
	defaultDynamodbTodosTableName = "todos"
)

type Settings struct {
	DynamodbTodosTableName string
}

func NewSettings() Settings {
	dynamodbTodosTableName := os.Getenv("DynamodbTodosTableName")
	if dynamodbTodosTableName == "" {
		dynamodbTodosTableName = defaultDynamodbTodosTableName
	}

	return Settings{
		DynamodbTodosTableName: dynamodbTodosTableName,
	}
}

func NewMockSettings() Settings {
	return Settings{
		DynamodbTodosTableName: defaultDynamodbTodosTableName,
	}
}
