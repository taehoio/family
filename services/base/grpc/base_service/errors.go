package base_service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	InvalidAccountIDError = status.Error(codes.InvalidArgument, "invalid account_id")
	PermissionDeniedError = status.Error(codes.PermissionDenied, "Forbidden")
)
