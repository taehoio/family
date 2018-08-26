package handlers

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/services/accounts/crypt"
	"github.com/taeho-io/family/services/accounts/repos/accounts_repo"
	"google.golang.org/grpc"
)

type LogInHandlerFunc func(ctx context.Context, req *accounts.LogInRequest) (*accounts.LogInResponse, error)

func LogIn(
	accountsTable *accounts_repo.Table,
	crypt crypt.IFace,
	authFunc func(ctx context.Context, in *auth.AuthRequest, opts ...grpc.CallOption) (*auth.AuthResponse, error),
) LogInHandlerFunc {
	return func(ctx context.Context, req *accounts.LogInRequest) (*accounts.LogInResponse, error) {
		account, err := accountsTable.GetByEmail(req.Email)
		if err != nil || account == nil || account.AccountID == "" {
			return nil, status.Error(codes.Unauthenticated, "")
		}

		acc, err := accountsTable.Get(account.AccountID)
		if err != nil || acc == nil {
			return nil, status.Error(codes.Unauthenticated, "")
		}

		isCorrectPassword := crypt.CheckHashedPassword(acc.HashedPassword, req.Password)
		if !isCorrectPassword {
			return nil, status.Error(codes.Unauthenticated, "")
		}

		authRes, err := authFunc(ctx, &auth.AuthRequest{
			AccountId: acc.AccountID,
		})
		if err != nil {
			return nil, err
		}

		return &accounts.LogInResponse{
			AccountId:    authRes.AccountId,
			AccessToken:  authRes.AccessToken,
			ExpiresIn:    authRes.ExpiresIn,
			RefreshToken: authRes.RefreshToken,
		}, nil
	}
}
