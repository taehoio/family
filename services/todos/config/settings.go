package config

import (
	"os"
)

const (
	serviceName                              = "todos"
	defaultDynamodbTodosTableName            = "todos"
	defaultDynamodbTodoGroupsTableName       = "todo_groups"
	defaultDynamodbTodoGroupPermitsTableName = "todo_group_permits"
)

type Settings struct {
	DynamodbTodosTableName            string
	DynamodbTodoGroupsTableName       string
	DynamodbTodoGroupPermitsTableName string
}

func NewSettings() Settings {
	dynamodbTodosTableName := os.Getenv("DynamodbTodosTableName")
	if dynamodbTodosTableName == "" {
		dynamodbTodosTableName = defaultDynamodbTodosTableName
	}

	dynamodbTodoGroupsTableName := os.Getenv("DynamodbTodoGroupsTableName")
	if dynamodbTodoGroupsTableName == "" {
		dynamodbTodoGroupsTableName = defaultDynamodbTodoGroupsTableName
	}

	dynamodbTodoGroupPermitsTableName := os.Getenv("DynamodbTodoGroupPermitsTableName")
	if dynamodbTodoGroupPermitsTableName == "" {
		dynamodbTodoGroupPermitsTableName = defaultDynamodbTodoGroupPermitsTableName
	}

	return Settings{
		DynamodbTodosTableName:            dynamodbTodosTableName,
		DynamodbTodoGroupsTableName:       dynamodbTodoGroupsTableName,
		DynamodbTodoGroupPermitsTableName: dynamodbTodoGroupPermitsTableName,
	}
}

func NewMockSettings() Settings {
	return Settings{
		DynamodbTodosTableName:            defaultDynamodbTodosTableName,
		DynamodbTodoGroupsTableName:       defaultDynamodbTodoGroupsTableName,
		DynamodbTodoGroupPermitsTableName: defaultDynamodbTodoGroupPermitsTableName,
	}
}
