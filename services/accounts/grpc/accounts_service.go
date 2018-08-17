package grpc

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
	"github.com/taeho-io/family/services/accounts/config"
	"github.com/taeho-io/family/services/accounts/crypt"
	"github.com/taeho-io/family/services/accounts/grpc/handlers"
	"github.com/taeho-io/family/services/accounts/repos/accounts_repo"
	"github.com/taeho-io/family/services/base/aws"
	"github.com/taeho-io/family/services/base/aws/dynamodb"
	grpcService "github.com/taeho-io/family/services/base/grpc"
)

type IFace interface {
	grpcService.IFace

	Config() config.IFace
	Crypt() crypt.IFace
	Dynamodb() dynamodb.IFace
	AccountsTable() *accounts_repo.Table
}

type Service struct {
	IFace

	cfg           config.IFace
	crypt         crypt.IFace
	ddb           dynamodb.IFace
	accountsTable *accounts_repo.Table
}

func New(cfg config.IFace) (*Service, error) {
	bcrypt := crypt.New()

	a, err := aws.New()
	if err != nil {
		return nil, err
	}
	ddb := dynamodb.New(a)

	return &Service{
		cfg:           cfg,
		crypt:         bcrypt,
		ddb:           ddb,
		accountsTable: accounts_repo.New(ddb, cfg),
	}, nil
}

func NewMock() (*Service, error) {
	return New(config.NewMock())
}

func (s *Service) RegisterService(srv *grpc.Server) {
	accounts.RegisterAccountsServiceServer(srv, s)
}

func (s *Service) Config() config.IFace {
	return s.cfg
}

func (s *Service) Crypt() crypt.IFace {
	return s.crypt
}

func (s *Service) Dynamodb() dynamodb.IFace {
	return s.ddb
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
