package grpc

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/taeho-io/family/idl/generated/go/pb/family/account"
	"github.com/taeho-io/family/svc/account/config"
	"github.com/taeho-io/family/svc/account/crypt"
	"github.com/taeho-io/family/svc/account/grpc/handlers"
	accountRepo "github.com/taeho-io/family/svc/account/repos/account"
	accountEmailRepo "github.com/taeho-io/family/svc/account/repos/account_email"
	"github.com/taeho-io/family/svc/srv/aws"
	"github.com/taeho-io/family/svc/srv/aws/dynamodb"
	srvGRPC "github.com/taeho-io/family/svc/srv/grpc"
)

type IFace interface {
	Config() config.IFace
	Crypt() crypt.IFace
	Dynamodb() dynamodb.IFace
	AccountTable() *accountRepo.Table
	AccountEmailTable() *accountEmailRepo.Table
}

type Service struct {
	srvGRPC.IFace

	cfg               config.IFace
	crypt             crypt.IFace
	ddb               dynamodb.IFace
	accountTable      *accountRepo.Table
	accountEmailTable *accountEmailRepo.Table
}

func New(cfg config.IFace) (*Service, error) {
	bcrypt := crypt.New()

	a, err := aws.New()
	if err != nil {
		return nil, err
	}
	ddb := dynamodb.New(a)

	return &Service{
		cfg:               cfg,
		crypt:             bcrypt,
		ddb:               ddb,
		accountTable:      accountRepo.New(ddb, cfg),
		accountEmailTable: accountEmailRepo.New(ddb, cfg),
	}, nil
}

func NewMock() (*Service, error) {
	return New(config.NewMock())
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

func (s *Service) AccountTable() *accountRepo.Table {
	return s.accountTable
}

func (s *Service) AccountEmailTable() *accountEmailRepo.Table {
	return s.accountEmailTable
}

func (s *Service) RegisterService(srv *grpc.Server) {
	account.RegisterAccountServiceServer(srv, s)
}

func (s *Service) Register(ctx context.Context, req *account.RegisterRequest) (*account.RegisterResponse, error) {
	return handlers.Register(s.AccountTable(), s.AccountEmailTable(), s.Crypt())(ctx, req)
}

func (s *Service) LogIn(ctx context.Context, req *account.LogInRequest) (*account.LogInResponse, error) {
	return handlers.LogIn(s.AccountTable(), s.AccountEmailTable(), s.Crypt())(ctx, req)
}
