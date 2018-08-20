package dynamodb_service

import (
	"github.com/taeho-io/family/services/base/aws"
	"github.com/taeho-io/family/services/base/aws/dynamodb"
	"github.com/taeho-io/family/services/base/config"
	"github.com/taeho-io/family/services/base/grpc/base_service"
)

type IFace interface {
	base_service.IFace

	Dynamodb() dynamodb.IFace
}

type Service struct {
	base_service.IFace

	ddb dynamodb.IFace
}

func New(cfg config.IFace) (IFace, error) {
	a, err := aws.New()
	if err != nil {
		return nil, err
	}
	ddb := dynamodb.New(a)

	svc := &Service{
		IFace: base_service.New(cfg),
		ddb:   ddb,
	}
	return svc, nil
}

func (s *Service) Dynamodb() dynamodb.IFace {
	return s.ddb
}
