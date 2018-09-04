package base

import (
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/services/discovery"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	RequestIDHeaderKey = "x-request-id"
	RequestIDKey       = "request_id"
	AccountIDHeaderKey = "x-account-id"
	AccountIDKey       = "account_id"
	ShouldAuthKey      = "should_auth"
	AuthBearerScheme   = "bearer"
)

func RequestIDUnaryServerInterceptor(
	ctx context.Context, req interface{},
	_ *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (
	interface{},
	error,
) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return handler(ctx, req)
	}

	var requestID string
	if ctx.Value(RequestIDKey) != nil {
		requestID = ctx.Value(RequestIDKey).(string)
	} else if requestIDs, ok := md[RequestIDHeaderKey]; ok && len(requestIDs) > 0 {
		requestID = requestIDs[0]
	} else {
		requestID = xid.New().String()
	}

	newCtx := context.WithValue(ctx, RequestIDKey, requestID)
	newCtx = ctxlogrus.ToContext(newCtx, logrus.WithFields(logrus.Fields{RequestIDKey: requestID}))
	newCtx = metadata.AppendToOutgoingContext(newCtx, RequestIDHeaderKey, requestID)
	return handler(newCtx, req)
}

func LogrusUnaryServerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (
	interface{},
	error,
) {
	entry := logrus.NewEntry(logrus.StandardLogger())

	ctxlogrus.AddFields(ctx, logrus.Fields{
		RequestIDKey: ctx.Value(RequestIDKey),
	})

	return grpc_logrus.UnaryServerInterceptor(entry)(ctx, req, info, handler)
}

func AuthWithWhiteListUnaryServerInterceptor(prefixes []string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		fullMethod := info.FullMethod
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
) (
	interface{},
	error,
) {
	return grpc_auth.UnaryServerInterceptor(authFunc)(ctx, req, info, handler)
}

func authFunc(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, nil
	}

	if accountIDs, ok := md[AccountIDHeaderKey]; ok && len(accountIDs) > 0 {
		accountID := accountIDs[0]
		newCtx := context.WithValue(ctx, AccountIDKey, accountID)
		newCtx = ctxlogrus.ToContext(newCtx, logrus.WithField(AccountIDKey, accountID))
		newCtx = metadata.AppendToOutgoingContext(newCtx, AccountIDHeaderKey, accountID)
		return newCtx, nil
	}

	shouldAuth := ctx.Value(ShouldAuthKey)
	if shouldAuth != nil && !shouldAuth.(bool) {
		return ctx, nil
	}

	token, err := grpc_auth.AuthFromMD(ctx, AuthBearerScheme)
	if err != nil {
		return nil, err
	}

	authSvc, err := discovery.NewAuthServiceClient()
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
	newCtx = ctxlogrus.ToContext(newCtx, logrus.WithField(AccountIDKey, accountID))
	newCtx = metadata.AppendToOutgoingContext(newCtx, AccountIDHeaderKey, accountID)
	return newCtx, nil
}

type GetAccountIDFromContextFunc func(context.Context) string
type HasPermissionByAccountIDFunc func(context.Context, string) error

func GetAccountIDFromContext(ctx context.Context) string {
	switch accountID := ctx.Value(AccountIDKey).(type) {
	case string:
		return accountID
	default:
		return ""
	}
}

func HasPermissionByAccountID(ctx context.Context, accountID string) error {
	if accountID == "" {
		return PermissionDeniedError
	}

	accountIDFromCtx := ctx.Value(AccountIDKey)
	if accountIDFromCtx != accountID {
		return PermissionDeniedError
	}

	return nil
}
