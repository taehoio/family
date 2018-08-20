package handlers

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	InvalidAccountIDError   = status.Error(codes.InvalidArgument, "invalid account_id")
	InvalidTodoGroupIDError = status.Error(codes.InvalidArgument, "invalid todo_group_id")
	InvalidTitleError       = status.Error(codes.InvalidArgument, "invalid title")
	PermissionDeniedError   = status.Error(codes.PermissionDenied, "Forbidden")
)
