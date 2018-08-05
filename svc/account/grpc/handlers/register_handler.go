package handlers

import (
	"github.com/pkg/errors"
	"go.uber.org/multierr"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"fmt"

	"github.com/rs/xid"
	"github.com/taeho-io/family/idl/generated/go/pb/family/account"
	"github.com/taeho-io/family/svc/account/crypt"
	accountRepo "github.com/taeho-io/family/svc/account/repos/account"
	accountEmailRepo "github.com/taeho-io/family/svc/account/repos/account_email"
)

type RegisterHandlerFunc func(ctx context.Context, req *account.RegisterRequest) (*account.RegisterResponse, error)

func Register(accountTable *accountRepo.Table, accountEmailTable *accountEmailRepo.Table, crypt crypt.IFace) RegisterHandlerFunc {
	return func(ctx context.Context, req *account.RegisterRequest) (*account.RegisterResponse, error) {
		req.AuthType = account.AuthType_EMAIL

		err := validateRegisterRequest(req)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		hashedPassword, err := crypt.HashPassword(req.Password)
		if err != nil {
			return nil, err
		}

		fmt.Println(accountEmailTable.Table().Name())

		accountID, err := accountEmailTable.GetAccountIDByEmail(req.Email)
		if accountID != "" {
			return nil, status.Error(codes.AlreadyExists, "email already exists")
		}

		accountID = xid.New().String()
		if err := accountTable.Put(&accountRepo.Account{
			AccountID:      accountID,
			Type:           account.AuthType_EMAIL.String(),
			Email:          req.Email,
			HashedPassword: hashedPassword,
			FullName:       req.FullName,
		}); err != nil {
			return nil, err
		}

		if err := accountEmailTable.Put(&accountEmailRepo.AccountEmail{
			Email:     req.Email,
			AccountID: accountID,
		}); err != nil {
			return nil, err
		}

		return &account.RegisterResponse{
			AccountId: accountID,
			AuthType:  account.AuthType_EMAIL,
		}, nil
	}
}

func validateRegisterRequest(req *account.RegisterRequest) error {
	var err error
	if req.AuthType == account.AuthType_NONE {
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
