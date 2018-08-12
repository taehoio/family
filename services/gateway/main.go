package main

import (
	"flag"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/xissy/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"

	accountsGateway "github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
	authGateway "github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	accountsConfig "github.com/taeho-io/family/services/accounts/config"
	accountsGRPC "github.com/taeho-io/family/services/accounts/grpc"
	authConfig "github.com/taeho-io/family/services/auth/config"
	authGRPC "github.com/taeho-io/family/services/auth/grpc"
)

const (
	RequestIdHeaderKey = "x-request-id"
)

var (
	grpcServerPort      = 9001
	grpcServerOrigin    = "localhost:" + strconv.Itoa(grpcServerPort)
	gatewayServerOrigin = ":" + os.Getenv("PORT")
)

var (
	authServerEndpoint = flag.String(
		"auth_server_endpoint",
		grpcServerOrigin,
		"endpoint of AuthServer",
	)
	accountServerEndpoint = flag.String(
		"account_server_endpoint",
		grpcServerOrigin,
		"endpoint of AccountServer",
	)
)

func requestIdMatcher(headerName string) (mdName string, ok bool) {
	lowerCasedHeaderName := strings.ToLower(headerName)
	if lowerCasedHeaderName == RequestIdHeaderKey {
		return lowerCasedHeaderName, true
	}
	return "", false
}

func serveGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(requestIdMatcher))
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := authGateway.RegisterAuthServiceHandlerFromEndpoint(
		ctx,
		mux,
		*authServerEndpoint,
		opts,
	); err != nil {
		return err
	}
	if err := accountsGateway.RegisterAccountsServiceHandlerFromEndpoint(
		ctx,
		mux,
		*accountServerEndpoint,
		opts,
	); err != nil {
		return err
	}

	return http.ListenAndServe(gatewayServerOrigin, mux)
}

func checkRequestIdUnaryServerInterceptor(
	ctx context.Context, req interface{},
	_ *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return handler(ctx, req)
	}
	if requestIds, ok := md[RequestIdHeaderKey]; ok && len(requestIds) > 0 {
		return handler(ctx, req)
	}
	requestIdUuid := uuid.NewV4()
	requestId := requestIdUuid.String()
	md[RequestIdHeaderKey] = []string{requestId}
	newCtx := metadata.NewOutgoingContext(ctx, md)
	return handler(newCtx, req)
}

func addRequestIdLogFieldUnaryServerInterceptor(
	ctx context.Context,
	req interface{},
	_ *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return handler(ctx, req)
	}
	if requestIds, ok := md[RequestIdHeaderKey]; ok && len(requestIds) > 0 {
		newCtx := ctxlogrus.ToContext(ctx, log.WithField("request_id", requestIds[0]))
		return handler(newCtx, req)
	}
	return handler(ctx, req)
}

func serveGRPCServers() error {
	lis, err := net.Listen("tcp", grpcServerOrigin)
	if err != nil {
		return err
	}

	logrusEntry := log.NewEntry(log.StandardLogger())
	var opts []grpc_logrus.Option

	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			checkRequestIdUnaryServerInterceptor,
			addRequestIdLogFieldUnaryServerInterceptor,
			grpc_logrus.UnaryServerInterceptor(logrusEntry, opts...),
		),
	)

	authCfg := authConfig.New(authConfig.NewSettings())
	authService := authGRPC.New(authCfg)
	authService.RegisterService(grpcServer)

	accountCfg := accountsConfig.New(accountsConfig.NewSettings())
	accountService, err := accountsGRPC.New(accountCfg)
	if err != nil {
		return err
	}
	accountService.RegisterService(grpcServer)

	// Register reflection server on gRPC server.
	reflection.Register(grpcServer)
	return grpcServer.Serve(lis)
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.ApexUpJSONFormatter{
		DisableTimestamp: true,
	})

	go func() {
		log.WithField("server_type", "grpc").Info("initializing")

		err := serveGRPCServers()
		if err != nil {
			log.WithField("server_type", "grpc").WithError(err).Fatal("failed to listen")
			return
		}
	}()

	log.WithField("server_type", "grpc_gw").Info("initializing")
	if err := serveGateway(); err != nil {
		log.WithField("server_type", "grpc_tw").WithError(err).Fatal("failed to listen")
		log.Fatalf("gateway: failed to listen: %x", err.Error())
	}
}
