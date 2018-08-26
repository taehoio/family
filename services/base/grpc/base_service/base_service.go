package base_service

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/taeho-io/family/services/base/config"
	"github.com/taeho-io/family/services/base/grpc/interceptors"
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
			"product_name": cfg.ProductName(),
			"service_name": cfg.ServiceName(),
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

type GetAccountIDFromContextFunc func(context.Context) string
type HasPermissionByAccountIDFunc func(context.Context, string) error

func GetAccountIDFromContext(ctx context.Context) string {
	return ctx.Value(interceptors.AccountIDKey).(string)
}

func HasPermissionByAccountID(ctx context.Context, accountID string) error {
	if accountID == "" {
		return InvalidAccountIDError
	}

	accountIDFromCtx := ctx.Value(interceptors.AccountIDKey)
	if accountIDFromCtx != accountID {
		return status.Error(codes.PermissionDenied, "Forbidden")
	}

	return nil
}
