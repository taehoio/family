package config

import "os"

const (
	namespace = "account"
)

type Settings struct {
	AwsDynamodbRegion string
}

func NewSettings() Settings {
	return Settings{
		AwsDynamodbRegion: os.Getenv("AWS_DEFAULT_REGION"),
	}
}

func NewMockSettings() Settings {
	return Settings{
		AwsDynamodbRegion: "us-west-2",
	}
}
