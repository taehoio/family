package todogroups

import (
	"os"
)

const (
	serviceName                              = "todogroups"
	defaultDynamodbTodoGroupsTableName       = "todo_groups"
	defaultDynamodbTodoGroupPermitsTableName = "todo_group_permits"
)

type Settings struct {
	DynamodbTodoGroupsTableName       string
	DynamodbTodoGroupPermitsTableName string
}

func NewSettings() Settings {
	dynamodbTodoGroupsTableName := os.Getenv("DynamodbTodoGroupsTableName")
	if dynamodbTodoGroupsTableName == "" {
		dynamodbTodoGroupsTableName = defaultDynamodbTodoGroupsTableName
	}

	dynamodbTodoGroupPermitsTableName := os.Getenv("DynamodbTodoGroupPermitsTableName")
	if dynamodbTodoGroupPermitsTableName == "" {
		dynamodbTodoGroupPermitsTableName = defaultDynamodbTodoGroupPermitsTableName
	}

	return Settings{
		DynamodbTodoGroupsTableName:       dynamodbTodoGroupsTableName,
		DynamodbTodoGroupPermitsTableName: dynamodbTodoGroupPermitsTableName,
	}
}

func NewMockSettings() Settings {
	return Settings{
		DynamodbTodoGroupsTableName:       defaultDynamodbTodoGroupsTableName,
		DynamodbTodoGroupPermitsTableName: defaultDynamodbTodoGroupPermitsTableName,
	}
}
