package base

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GrpcService interface {
	Config() Config
	Logger() *logrus.Entry

	RegisterService(*grpc.Server)
}

type DefaultGrpcService struct {
	GrpcService

	cfg    Config
	logger *logrus.Entry
}

func NewGrpcService(cfg Config) GrpcService {
	svc := &DefaultGrpcService{
		cfg: cfg,
		logger: logrus.WithFields(logrus.Fields{
			"product_name": cfg.ProductName(),
			"service_name": cfg.ServiceName(),
		}),
	}
	return svc
}

func (s *DefaultGrpcService) Config() Config {
	return s.cfg
}

func (s *DefaultGrpcService) Logger() *logrus.Entry {
	return s.logger
}

func (s *DefaultGrpcService) Serve() error {
	return NotImplementedError
}
