package auth_service

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/idl/generated/go/pb/family/discovery"
	"github.com/taeho-io/family/services/auth/config"
	"github.com/taeho-io/family/services/auth/grpc/auth_service/handlers"
	"github.com/taeho-io/family/services/auth/token"
	"github.com/taeho-io/family/services/base/grpc/base_service"
	"github.com/taeho-io/family/services/base/grpc/interceptors"
	"github.com/taeho-io/family/services/discovery/grpc/discovery_service"
)

type IFace interface {
	base_service.IFace

	Token() token.IFace
}

type Service struct {
	base_service.IFace

	tkn token.IFace
}

func New(cfg config.IFace) (s *Service) {
	return &Service{
		IFace: base_service.New(cfg),
		tkn:   token.New(cfg),
	}
}

func NewMock() (s *Service) {
	return New(config.NewMock())
}

func (s *Service) RegisterService(srv *grpc.Server) {
	auth.RegisterAuthServiceServer(srv, s)
}

func (s *Service) Token() token.IFace {
	return s.tkn
}

func (s *Service) Auth(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {
	return handlers.Auth(s.Config().(config.IFace), s.Token())(ctx, req)
}

func (s *Service) Validate(ctx context.Context, req *auth.ValidateRequest) (*auth.ValidateResponse, error) {
	return handlers.Validate(s.Token())(ctx, req)
}

func (s *Service) Refresh(ctx context.Context, req *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	return handlers.Refresh(s.Config().(config.IFace), s.Token())(ctx, req)
}

func (s *Service) Parse(ctx context.Context, req *auth.ParseRequest) (*auth.ParseResponse, error) {
	return handlers.Parse(s.Token())(ctx, req)
}

func Serve() error {
	addr := discovery_service.ServiceAddrMap[discovery.Service_AUTH]
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	whiteListPrefixes := []string{
		"/",
	}

	svr := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			interceptors.RequestIDUnaryServerInterceptor,
			interceptors.LogrusUnaryServerInterceptor,
			interceptors.AuthWithWhiteListUnaryServerInterceptor(whiteListPrefixes),
			interceptors.AuthUnaryServerInterceptor,
		),
	)

	svc := New(config.New(config.NewSettings()))

	svc.RegisterService(svr)
	reflection.Register(svr)
	return svr.Serve(lis)
}
