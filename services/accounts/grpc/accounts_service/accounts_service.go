package accounts_service

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
	"github.com/taeho-io/family/idl/generated/go/pb/family/discovery"
	"github.com/taeho-io/family/services/accounts/config"
	"github.com/taeho-io/family/services/accounts/crypt"
	"github.com/taeho-io/family/services/accounts/grpc/accounts_service/handlers"
	"github.com/taeho-io/family/services/accounts/repos/accounts_repo"
	"github.com/taeho-io/family/services/base/grpc/dynamodb_service"
	"github.com/taeho-io/family/services/base/grpc/interceptors"
	"github.com/taeho-io/family/services/discovery/grpc/discovery_service"
)

type IFace interface {
	dynamodb_service.IFace

	Crypt() crypt.IFace
	AccountsTable() *accounts_repo.Table
}

type Service struct {
	dynamodb_service.IFace

	crypt         crypt.IFace
	accountsTable *accounts_repo.Table
}

func New(cfg config.IFace) (IFace, error) {
	bcrypt := crypt.New()

	dynamodbSvc, err := dynamodb_service.New(cfg)
	if err != nil {
		return nil, err
	}

	return &Service{
		IFace:         dynamodbSvc,
		crypt:         bcrypt,
		accountsTable: accounts_repo.New(dynamodbSvc.Dynamodb(), cfg),
	}, nil
}

func NewMock() (IFace, error) {
	return New(config.NewMock())
}

func (s *Service) RegisterService(srv *grpc.Server) {
	accounts.RegisterAccountsServiceServer(srv, s)
}

func (s *Service) Crypt() crypt.IFace {
	return s.crypt
}

func (s *Service) AccountsTable() *accounts_repo.Table {
	return s.accountsTable
}

func (s *Service) Register(ctx context.Context, req *accounts.RegisterRequest) (*accounts.RegisterResponse, error) {
	return handlers.Register(s.AccountsTable(), s.Crypt())(ctx, req)
}

func (s *Service) LogIn(ctx context.Context, req *accounts.LogInRequest) (*accounts.LogInResponse, error) {
	return handlers.LogIn(s.AccountsTable(), s.Crypt())(ctx, req)
}

func Serve() error {
	addr := discovery_service.ServiceAddrMap[discovery.Service_ACCOUNTS]
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

	svc, err := New(config.New(config.NewSettings()))
	if err != nil {
		return err
	}

	svc.RegisterService(svr)
	reflection.Register(svr)
	return svr.Serve(lis)
}
