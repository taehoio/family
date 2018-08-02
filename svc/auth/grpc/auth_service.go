package auth

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/svc/auth/config"
	"github.com/taeho-io/family/svc/auth/grpc/handlers"
	"github.com/taeho-io/family/svc/auth/token"
)

type Service struct {
	config config.IFace

	tokenSrv token.Token
}

func New(cfg config.IFace, tokenSrv token.Token) (s *Service) {
	return &Service{
		config:   cfg,
		tokenSrv: tokenSrv,
	}
}

func (s *Service) Name() string {
	return s.config.Namespace()
}

func (s *Service) RegisterService(srv *grpc.Server) {
	auth.RegisterAuthServiceServer(srv, s)
}

func (s *Service) Auth(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {
	return handlers.Auth(s.config, s.tokenSrv)(ctx, req)
}

func (s *Service) Validate(ctx context.Context, req *auth.ValidateRequest) (*auth.ValidateResponse, error) {
	return handlers.Validate(s.tokenSrv)(ctx, req)
}

func (s *Service) Refresh(ctx context.Context, req *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	return handlers.Refresh(s.config, s.tokenSrv)(ctx, req)
}

func (s *Service) Parse(ctx context.Context, req *auth.ParseRequest) (*auth.ParseResponse, error) {
	return handlers.Parse(s.tokenSrv)(ctx, req)
}
