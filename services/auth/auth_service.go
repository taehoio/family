package auth

import (
	"net"

	"github.com/taeho-io/family/services/auth/internal/handler"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/idl/generated/go/pb/family/discovery"
	"github.com/taeho-io/family/services/auth/pkg/token"
	"github.com/taeho-io/family/services/base"
	discoveryService "github.com/taeho-io/family/services/discovery"
)

type Service interface {
	base.GrpcService
	auth.AuthServiceServer

	Token() token.Token
}

type DefaultService struct {
	base.GrpcService

	tkn token.Token
}

func NewService(cfg Config) Service {
	tokenCfg := token.NewConfig(
		cfg.Settings().TokenIssuer,
		cfg.Settings().AccessTokenExpiringDuration,
		cfg.Settings().RefreshTokenExpiringDuration,
	)

	return &DefaultService{
		GrpcService: base.NewGrpcService(cfg),
		tkn:         token.New(tokenCfg),
	}
}

func NewMockService() Service {
	return NewService(NewMockConfig())
}

func (s *DefaultService) RegisterService(srv *grpc.Server) {
	auth.RegisterAuthServiceServer(srv, s)
}

func (s *DefaultService) Token() token.Token {
	return s.tkn
}

func (s *DefaultService) Auth(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {
	return handler.Auth(s.Config().(Config).Settings().AccessTokenExpiringDuration, s.Token())(ctx, req)
}

func (s *DefaultService) Validate(ctx context.Context, req *auth.ValidateRequest) (*auth.ValidateResponse, error) {
	return handler.Validate(s.Token())(ctx, req)
}

func (s *DefaultService) Refresh(ctx context.Context, req *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	return handler.Refresh(s.Config().(Config).Settings().AccessTokenExpiringDuration, s.Token())(ctx, req)
}

func (s *DefaultService) Parse(ctx context.Context, req *auth.ParseRequest) (*auth.ParseResponse, error) {
	return handler.Parse(s.Token())(ctx, req)
}

func Serve() error {
	addr := discoveryService.ServiceAddrMap[discovery.Service_AUTH]
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	whiteListPrefixes := []string{
		"/",
	}

	svr := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			base.RequestIDUnaryServerInterceptor,
			base.LogrusUnaryServerInterceptor,
			base.AuthWithWhiteListUnaryServerInterceptor(whiteListPrefixes),
			base.AuthUnaryServerInterceptor,
		),
	)

	svc := NewService(NewConfig(NewSettings()))

	svc.RegisterService(svr)
	reflection.Register(svr)
	return svr.Serve(lis)
}
