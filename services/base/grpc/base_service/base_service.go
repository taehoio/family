package base_service

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/taeho-io/family/services/base/config"
)

type IFace interface {
	Config() config.IFace
	Logger() *logrus.Entry
	RegisterService(*grpc.Server)
}

type Service struct {
	IFace

	cfg    config.IFace
	logger *logrus.Entry
}

func New(cfg config.IFace) IFace {
	svc := &Service{
		cfg: cfg,
		logger: logrus.WithFields(logrus.Fields{
			"productName": cfg.ProductName(),
			"serviceName": cfg.ServiceName(),
		}),
	}
	return svc
}

func (s *Service) Config() config.IFace {
	return s.cfg
}

func (s *Service) Logger() *logrus.Entry {
	return s.logger
}
