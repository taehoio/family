package handler

import (
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/services/accounts/internal/repo"
	"github.com/taeho-io/family/services/accounts/pkg/crypt"
)

type LogInHandlerFunc func(ctx context.Context, req *accounts.LogInRequest) (*accounts.LogInResponse, error)

func LogIn(
	accountsRepo repo.AccountsRepo,
	crypt crypt.Crypt,
	authServiceClient auth.AuthServiceClient,
) LogInHandlerFunc {
	return func(ctx context.Context, req *accounts.LogInRequest) (*accounts.LogInResponse, error) {
		account, err := accountsRepo.GetByEmail(req.Email)
		if err != nil || account == nil || account.AccountID == "" {
			return nil, UnauthorizedError
		}

		acc, err := accountsRepo.GetByID(account.AccountID)
		if err != nil || acc == nil {
			return nil, UnauthorizedError
		}

		isCorrectPassword := crypt.CheckHashedPassword(acc.HashedPassword, req.Password)
		if !isCorrectPassword {
			return nil, UnauthorizedError
		}

		authRes, err := authServiceClient.Auth(ctx, &auth.AuthRequest{
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
