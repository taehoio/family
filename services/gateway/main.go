package main

import (
	"flag"
	"net/http"
	"os"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	log "github.com/sirupsen/logrus"
	"github.com/xissy/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/idl/generated/go/pb/family/discovery"
	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
	"github.com/taeho-io/family/services/accounts/grpc/accounts_service"
	"github.com/taeho-io/family/services/auth/grpc/auth_service"
	"github.com/taeho-io/family/services/discovery/grpc/discovery_service"
	"github.com/taeho-io/family/services/todo_groups/grpc/todo_groups_service"
)

var (
	gatewayAddr        = ":" + os.Getenv("PORT")
	requestIdHeaderKey = "x-request-id"

	authServerEndpoint = flag.String(
		"auth_server_endpoint",
		discovery_service.ServiceAddrMap[discovery.Service_AUTH],
		"endpoint of AuthServer",
	)
	accountsServerEndpoint = flag.String(
		"accounts_server_endpoint",
		discovery_service.ServiceAddrMap[discovery.Service_ACCOUNTS],
		"endpoint of AccountsServer",
	)
	todoGroupsServerEndpoint = flag.String(
		"todo_groups_server_endpoint",
		discovery_service.ServiceAddrMap[discovery.Service_TODOGROUPS],
		"endpoint of TodoGroupsServer",
	)
)

func requestIDMatcher(headerName string) (mdName string, ok bool) {
	lowerCasedHeaderName := strings.ToLower(headerName)
	if lowerCasedHeaderName == requestIdHeaderKey {
		return lowerCasedHeaderName, true
	}
	return "", false
}

func serveGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(requestIDMatcher))
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := auth.RegisterAuthServiceHandlerFromEndpoint(
		ctx,
		mux,
		*authServerEndpoint,
		opts,
	); err != nil {
		return err
	}

	if err := accounts.RegisterAccountsServiceHandlerFromEndpoint(
		ctx,
		mux,
		*accountsServerEndpoint,
		opts,
	); err != nil {
		return err
	}

	if err := todo_groups.RegisterTodoGroupsServiceHandlerFromEndpoint(
		ctx,
		mux,
		*todoGroupsServerEndpoint,
		opts,
	); err != nil {
		return err
	}

	return http.ListenAndServe(gatewayAddr, mux)
}

func startGRPCServices() error {
	errs := make(chan error, 1)

	type serveFunc func() error

	serveFuncs := []serveFunc{
		auth_service.Serve,
		accounts_service.Serve,
		todo_groups_service.Serve,
	}

	for _, serve := range serveFuncs {
		go func(serve serveFunc) {
			if err := serve(); err != nil {
				errs <- err
			}
		}(serve)
	}

	if err, open := <-errs; open {
		return err
	}

	return nil
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.ApexUpJSONFormatter{
		DisableTimestamp: true,
	})

	go func() {
		log.WithField("server_type", "grpc").Info("initializing")

		err := startGRPCServices()
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
