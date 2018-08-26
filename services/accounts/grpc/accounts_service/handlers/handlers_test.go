package handlers

import (
	"os"
	"testing"

	"github.com/icrowley/fake"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/services/accounts/crypt"
	"github.com/taeho-io/family/services/accounts/repos/accounts_repo"
)

var (
	accountsTableMock *accounts_repo.Table
	cryptMock         crypt.IFace
	authMock          func(context.Context, *auth.AuthRequest, ...grpc.CallOption) (*auth.AuthResponse, error)

	register func(context.Context, *accounts.RegisterRequest) (*accounts.RegisterResponse, error)
	logIn    func(context.Context, *accounts.LogInRequest) (*accounts.LogInResponse, error)

	testAuthTypeEmail = accounts.AuthType_EMAIL
	testFullName      = fake.FullName()
	testEmail         = fake.EmailAddress()
	testPassword      = fake.SimplePassword()
)

func TestMain(m *testing.M) {
	accountsTableMock = accounts_repo.NewMock()
	cryptMock = crypt.New()
	authMock = func(_ context.Context, req *auth.AuthRequest, _ ...grpc.CallOption) (*auth.AuthResponse, error) {
		return &auth.AuthResponse{
			AccountId:    req.AccountId,
			AccessToken:  "test_access_token",
			RefreshToken: "test_refresh_token",
			ExpiresIn:    3600,
		}, nil
	}

	register = Register(accountsTableMock, cryptMock)
	logIn = LogIn(accountsTableMock, cryptMock, authMock)

	retCode := m.Run()
	os.Exit(retCode)
}
