package handlers

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/taeho-io/family/idl/generated/go/pb/family/account"
	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/svc/account/crypt"
	accountRepo "github.com/taeho-io/family/svc/account/repos/account"
	accountEmailRepo "github.com/taeho-io/family/svc/account/repos/account_email"
	"github.com/taeho-io/family/svc/discovery"
)

type LogInHandlerFunc func(ctx context.Context, req *account.LogInRequest) (*account.LogInResponse, error)

func LogIn(accountTable *accountRepo.Table, accountEmailTable *accountEmailRepo.Table, crypt crypt.IFace) LogInHandlerFunc {
	return func(ctx context.Context, req *account.LogInRequest) (*account.LogInResponse, error) {
		accountID, err := accountEmailTable.GetAccountIDByEmail(req.Email)
		if err != nil || accountID == "" {
			return nil, status.Error(codes.Unauthenticated, "")
		}
		acc, err := accountTable.Get(accountID)
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
		return &account.LogInResponse{
			AccountId:    authResponse.AccountId,
			AccessToken:  authResponse.AccessToken,
			ExpiresIn:    authResponse.ExpiresIn,
			RefreshToken: authResponse.RefreshToken,
		}, nil
	}
}
