package handler

import (
	"time"

	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/services/auth/pkg/jwt"
	"github.com/taeho-io/family/services/auth/pkg/token"
)

type RefreshHandlerFunc func(context.Context, *auth.RefreshRequest) (*auth.RefreshResponse, error)

func Refresh(accessTokenExpiringDuration time.Duration, tkn token.Token) RefreshHandlerFunc {
	return func(ctx context.Context, req *auth.RefreshRequest) (*auth.RefreshResponse, error) {
		claims, err := jwt.ParseToken(req.RefreshToken)
		if err != nil {
			return nil, err
		}

		accessTokenString, err := tkn.NewAccessToken(claims.Audience)
		if err != nil {
			return nil, err
		}

		return &auth.RefreshResponse{
			AccessToken: accessTokenString,
			ExpiresIn:   int64(accessTokenExpiringDuration.Seconds()),
		}, nil
	}
}
