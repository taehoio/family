package handlers

import (
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/services/auth/token"
)

type ParseHandlerFunc func(context.Context, *auth.ParseRequest) (*auth.ParseResponse, error)

func Parse(tkn token.IFace) ParseHandlerFunc {
	return func(ctx context.Context, req *auth.ParseRequest) (*auth.ParseResponse, error) {
		claims, err := tkn.ParseToken(req.AccessToken)
		if err != nil {
			return nil, err
		}

		return &auth.ParseResponse{
			AccountId: claims.Audience,
		}, nil
	}
}
