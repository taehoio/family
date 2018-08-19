package interceptors

import (
	"fmt"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/services/discovery/grpc/discovery_service"
)

const (
	RequestIDHeaderKey = "x-request-id"
	RequestIDKey       = "request_id"
	AccountIDKey       = "account_id"
	ShouldAuthKey      = "should_auth"
	AuthBearerScheme   = "bearer"
)

func NewRequestIdIfNotExistsUnaryServerInterceptor(
	ctx context.Context, req interface{},
	_ *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return handler(ctx, req)
	}

	if requestIds, ok := md[RequestIDHeaderKey]; ok && len(requestIds) > 0 {
		return handler(ctx, req)
	}

	requestIdUuid := uuid.NewV4()
	requestId := requestIdUuid.String()
	md[RequestIDHeaderKey] = []string{requestId}
	newCtx := metadata.NewOutgoingContext(ctx, md)
	return handler(newCtx, req)
}

func ForwardRequestIdLogFieldUnaryServerInterceptor(
	ctx context.Context,
	req interface{},
	_ *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return handler(ctx, req)
	}

	if requestIds, ok := md[RequestIDHeaderKey]; ok && len(requestIds) > 0 {
		newCtx := ctxlogrus.ToContext(ctx, logrus.WithField(RequestIDKey, requestIds[0]))
		return handler(newCtx, req)
	}

	return handler(ctx, req)
}

func LogrusUnaryServerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	entry := logrus.NewEntry(logrus.StandardLogger())

	return grpc_logrus.UnaryServerInterceptor(entry)(ctx, req, info, handler)
}

func AuthWhileListUnaryServerInterceptor(prefixes []string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		fullMethod := info.FullMethod
		fmt.Println(fullMethod)
		for _, prefix := range prefixes {
			if strings.HasPrefix(fullMethod, prefix) {
				newCtx := context.WithValue(ctx, ShouldAuthKey, false)
				return handler(newCtx, req)
			}
		}

		newCtx := context.WithValue(ctx, ShouldAuthKey, true)
		return handler(newCtx, req)
	}
}

func AuthUnaryServerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	authFunc := func(ctx context.Context) (context.Context, error) {
		shouldAuth := ctx.Value(ShouldAuthKey)
		if !shouldAuth.(bool) {
			return ctx, nil
		}

		if ctx.Value(AccountIDKey) != nil {
			return ctx, nil
		}

		token, err := grpc_auth.AuthFromMD(ctx, AuthBearerScheme)
		if err != nil {
			return nil, err
		}

		authSvc, err := discovery_service.NewAuthServiceClient()
		if err != nil {
			return ctx, err
		}

		parseResponse, err := authSvc.Parse(ctx, &auth.ParseRequest{
			AccessToken: token,
		})
		if err != nil {
			return nil, err
		}

		accountID := parseResponse.GetAccountId()
		newCtx := context.WithValue(ctx, AccountIDKey, accountID)
		return newCtx, nil
	}

	return grpc_auth.UnaryServerInterceptor(authFunc)(ctx, req, info, handler)
}
