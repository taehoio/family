package handlers

import (
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/svc/auth/token"
)

type ValidateHandlerFunc func(context.Context, *auth.ValidateRequest) (*auth.ValidateResponse, error)

func Validate(tokenSrv token.Token) ValidateHandlerFunc {
	return func(ctx context.Context, req *auth.ValidateRequest) (*auth.ValidateResponse, error) {
		err := tokenSrv.ValidateToken(req.AccessToken)
		isValid := err != nil
		return &auth.ValidateResponse{
			IsValid: isValid,
		}, nil
	}
}
