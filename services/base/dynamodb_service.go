package base

import "strings"

type DynamodbService interface {
	GrpcService

	Dynamodb() Dynamodb
}

type DefaultDynamodbService struct {
	GrpcService

	ddb Dynamodb
}

func NewDynamodbService(cfg Config) (DynamodbService, error) {
	aws, err := NewAws()
	if err != nil {
		return nil, err
	}
	ddb := NewDynamodb(aws)

	svc := &DefaultDynamodbService{
		GrpcService: NewGrpcService(cfg),
		ddb:         ddb,
	}
	return svc, nil
}

func (s *DefaultDynamodbService) Dynamodb() Dynamodb {
	return s.ddb
}

func FullDynamodbTableName(cfg Config, tableName string) string {
	prefix := cfg.Prefix()
	return strings.Join([]string{prefix, tableName}, "-")
}
