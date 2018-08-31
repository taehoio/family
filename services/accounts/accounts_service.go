package accounts

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/idl/generated/go/pb/family/discovery"
	"github.com/taeho-io/family/services/accounts/internal/handler"
	"github.com/taeho-io/family/services/accounts/internal/repo"
	"github.com/taeho-io/family/services/accounts/pkg/crypt"
	"github.com/taeho-io/family/services/base"
	discoveryService "github.com/taeho-io/family/services/discovery"
)

type Service interface {
	base.DynamodbService
	accounts.AccountsServiceServer

	Crypt() crypt.Crypt
	AccountsRepo() repo.AccountsRepo
	AuthServiceClient() auth.AuthServiceClient
}

type defaultService struct {
	base.DynamodbService

	crypt             crypt.Crypt
	accountsRepo      repo.AccountsRepo
	authServiceClient auth.AuthServiceClient
}

func NewService(cfg Config) (Service, error) {
	bcrypt := crypt.New(crypt.NewConfig())

	dynamodbSvc, err := base.NewDynamodbService(cfg)
	if err != nil {
		return nil, err
	}

	authServiceClient, err := discoveryService.NewAuthServiceClient()
	if err != nil {
		return nil, err
	}

	accountsRepo := repo.NewAccountsRepo(
		dynamodbSvc.Dynamodb(),
		repo.NewAccountsRepoConfig(
			base.FullDynamodbTableName(cfg, cfg.Settings().DynamodbAccountsTableName),
		),
	)

	return &defaultService{
		DynamodbService:   dynamodbSvc,
		crypt:             bcrypt,
		accountsRepo:      accountsRepo,
		authServiceClient: authServiceClient,
	}, nil
}

func NewMockService() (Service, error) {
	return NewService(NewMockConfig())
}

func (s *defaultService) RegisterService(srv *grpc.Server) {
	accounts.RegisterAccountsServiceServer(srv, s)
}

func (s *defaultService) Crypt() crypt.Crypt {
	return s.crypt
}

func (s *defaultService) AccountsRepo() repo.AccountsRepo {
	return s.accountsRepo
}

func (s *defaultService) AuthServiceClient() auth.AuthServiceClient {
	return s.authServiceClient
}

func (s *defaultService) Register(ctx context.Context, req *accounts.RegisterRequest) (*accounts.RegisterResponse, error) {
	return handler.Register(s.AccountsRepo(), s.Crypt())(ctx, req)
}

func (s *defaultService) LogIn(ctx context.Context, req *accounts.LogInRequest) (*accounts.LogInResponse, error) {
	return handler.LogIn(s.AccountsRepo(), s.Crypt(), s.AuthServiceClient())(ctx, req)
}

func Serve() error {
	addr := discoveryService.ServiceAddrMap[discovery.Service_ACCOUNTS]
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

	svc, err := NewService(NewConfig(NewSettings()))
	if err != nil {
		return err
	}

	svc.RegisterService(svr)
	reflection.Register(svr)
	return svr.Serve(lis)
}
