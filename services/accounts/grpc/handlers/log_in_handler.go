package handlers

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/services/accounts/crypt"
	"github.com/taeho-io/family/services/accounts/repos/accounts_repo"
	"github.com/taeho-io/family/services/discovery"
)

type LogInHandlerFunc func(ctx context.Context, req *accounts.LogInRequest) (*accounts.LogInResponse, error)

func LogIn(accountTable *accounts_repo.Table, crypt crypt.IFace) LogInHandlerFunc {
	return func(ctx context.Context, req *accounts.LogInRequest) (*accounts.LogInResponse, error) {
		account, err := accountTable.GetByEmail(req.Email)
		if err != nil || account == nil || account.AccountID == "" {
			return nil, status.Error(codes.Unauthenticated, "")
		}

		acc, err := accountTable.Get(account.AccountID)
		if err != nil || acc == nil {
			return nil, status.Error(codes.Unauthenticated, "")
		}

		isCorrectPassword := crypt.CheckHashedPassword(acc.HashedPassword, req.Password)
		if !isCorrectPassword {
			return nil, status.Error(codes.Unauthenticated, "")
		}

		authClient, err := discovery.NewAuthServiceClient()
		if err != nil {
			return nil, err
		}

		if md, ok := metadata.FromIncomingContext(ctx); ok {
			ctx = metadata.NewOutgoingContext(ctx, md)
		}

		authResponse, err := authClient.Auth(ctx, &auth.AuthRequest{
			AccountId: acc.AccountID,
		})
		if err != nil {
			return nil, err
		}

		return &accounts.LogInResponse{
			AccountId:    authResponse.AccountId,
			AccessToken:  authResponse.AccessToken,
			ExpiresIn:    authResponse.ExpiresIn,
			RefreshToken: authResponse.RefreshToken,
		}, nil
	}
}
