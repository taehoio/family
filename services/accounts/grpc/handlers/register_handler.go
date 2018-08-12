package handlers

import (
	"time"

	"github.com/pkg/errors"
	"github.com/rs/xid"
	"go.uber.org/multierr"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
	"github.com/taeho-io/family/services/accounts/crypt"
	"github.com/taeho-io/family/services/accounts/models"
	"github.com/taeho-io/family/services/accounts/repos/accounts_repo"
)

type RegisterHandlerFunc func(ctx context.Context, req *accounts.RegisterRequest) (*accounts.RegisterResponse, error)

func Register(accountTable *accounts_repo.Table, crypt crypt.IFace) RegisterHandlerFunc {
	return func(ctx context.Context, req *accounts.RegisterRequest) (*accounts.RegisterResponse, error) {
		req.AuthType = accounts.AuthType_EMAIL

		err := validateRegisterRequest(req)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		hashedPassword, err := crypt.HashPassword(req.Password)
		if err != nil {
			return nil, err
		}

		account, err := accountTable.GetByEmail(req.Email)
		if account != nil {
			return nil, status.Error(codes.AlreadyExists, "email already exists")
		}

		accountID := xid.New().String()
		currTime := time.Now()
		if err := accountTable.Put(&models.Account{
			AccountID:      accountID,
			Type:           accounts.AuthType_EMAIL.String(),
			Email:          req.Email,
			HashedPassword: hashedPassword,
			FullName:       req.FullName,
			CreateAt:       currTime,
			UpdatedAt:      currTime,
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
	var err error
	if req.AuthType == accounts.AuthType_NONE {
		err = multierr.Append(err, errors.New("invalid auth_type"))
	}
	if req.Email == "" {
		err = multierr.Append(err, errors.New("empty email"))
	}
	if req.Password == "" {
		err = multierr.Append(err, errors.New("empty password"))
	}
	return err
}
