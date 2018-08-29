package handler

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	InvalidTodoIDError     = status.Error(codes.InvalidArgument, "invalid todo_id")
	InvalidTitleError      = status.Error(codes.InvalidArgument, "invalid title")
	InvalidParentTypeError = status.Error(codes.InvalidArgument, "invalid parent_type")
	InvalidParentIDError   = status.Error(codes.InvalidArgument, "invalid parent_id")
)
