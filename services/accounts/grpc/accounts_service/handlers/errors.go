package handlers

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	UnauthorizedError       = status.Error(codes.Unauthenticated, "")
	InvalidAuthTypeError    = status.Error(codes.InvalidArgument, "invalid auth_type")
	InvalidFullNameError    = status.Error(codes.InvalidArgument, "invalid full_name")
	InvalidEmailError       = status.Error(codes.InvalidArgument, "invalid email")
	InvalidPasswordError    = status.Error(codes.InvalidArgument, "invalid password")
	EmailAlreadyExistsError = status.Error(codes.AlreadyExists, "email already exists")
)
