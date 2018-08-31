package handler

import (
	"time"

	"github.com/rs/xid"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
	"github.com/taeho-io/family/services/accounts/internal/model"
	"github.com/taeho-io/family/services/accounts/internal/repo"
	"github.com/taeho-io/family/services/accounts/pkg/crypt"
)

type RegisterHandlerFunc func(ctx context.Context, req *accounts.RegisterRequest) (*accounts.RegisterResponse, error)

func Register(accountsRepo repo.AccountsRepo, crypt crypt.Crypt) RegisterHandlerFunc {
	return func(ctx context.Context, req *accounts.RegisterRequest) (*accounts.RegisterResponse, error) {
		err := validateRegisterRequest(req)
		if err != nil {
			return nil, err
		}

		hashedPassword, err := crypt.HashPassword(req.Password)
		if err != nil {
			return nil, err
		}

		account, err := accountsRepo.GetByEmail(req.Email)
		if err != nil {
			return nil, err
		}
		if account != nil {
			return nil, EmailAlreadyExistsError
		}

		accountID := xid.New().String()
		now := time.Now()
		if err := accountsRepo.Put(&model.Account{
			AccountID:      accountID,
			Type:           accounts.AuthType_EMAIL,
			Email:          req.Email,
			HashedPassword: hashedPassword,
			FullName:       req.FullName,
			CreateAt:       now,
			UpdatedAt:      now,
		}); err != nil {
			return nil, err
		}

		return &accounts.RegisterResponse{
			AccountId: accountID,
			AuthType:  accounts.AuthType_EMAIL,
		}, nil
	}
}

func validateRegisterRequest(req *accounts.RegisterRequest) error {
	if req.AuthType == accounts.AuthType_NONE {
		return InvalidAuthTypeError
	}
	if req.FullName == "" {
		return InvalidFullNameError
	}
	if req.Email == "" {
		return InvalidEmailError
	}
	if req.Password == "" {
		return InvalidPasswordError
	}

	return nil
}
