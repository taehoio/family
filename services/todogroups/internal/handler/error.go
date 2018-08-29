package handler

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	InvalidTodoGroupIDError = status.Error(codes.InvalidArgument, "invalid todo_group_id")
	InvalidTitleError       = status.Error(codes.InvalidArgument, "invalid title")
)
