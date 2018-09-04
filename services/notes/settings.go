package notes

import "os"

const (
	serviceName                   = "notes"
	defaultDynamodbNotesTableName = "notes"
)

type Settings struct {
	DynamodbNotesTableName string
}

func NewSettings() Settings {
	dynamodbNotesTableName := os.Getenv("DynamodbNotesTableName")
	if dynamodbNotesTableName == "" {
		dynamodbNotesTableName = defaultDynamodbNotesTableName
	}

	return Settings{
		DynamodbNotesTableName: dynamodbNotesTableName,
	}
}

func NewMockSettings() Settings {
	return Settings{
		DynamodbNotesTableName: defaultDynamodbNotesTableName,
	}
}
