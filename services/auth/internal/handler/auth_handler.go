package handler

import (
	"time"

	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/services/auth/pkg/token"
)

type AuthHandlerFunc func(context.Context, *auth.AuthRequest) (*auth.AuthResponse, error)

func Auth(accessTokenExpiringDuration time.Duration, tkn token.Token) AuthHandlerFunc {
	return func(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {
		accessTokenString, err := tkn.NewAccessToken(req.AccountId)
		if err != nil {
			return nil, err
		}

		refreshTokenString, err := tkn.NewRefreshToken(req.AccountId)
		if err != nil {
			return nil, err
		}

		return &auth.AuthResponse{
			AccountId:    req.AccountId,
			AccessToken:  accessTokenString,
			RefreshToken: refreshTokenString,
			ExpiresIn:    int64(accessTokenExpiringDuration.Seconds()),
		}, nil
	}
}
