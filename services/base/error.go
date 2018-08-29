package base

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	NotImplementedError   = fmt.Errorf("not imeplemented")
	InvalidAccountIDError = status.Error(codes.InvalidArgument, "invalid account_id")
	PermissionDeniedError = status.Error(codes.PermissionDenied, "Forbidden")
)
