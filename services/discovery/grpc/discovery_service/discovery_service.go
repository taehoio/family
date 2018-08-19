package discovery_service

import (
	"google.golang.org/grpc"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/idl/generated/go/pb/family/discovery"
	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
)

var ServiceAddrMap = map[discovery.Service]string{
	discovery.Service_AUTH:       "localhost:9000",
	discovery.Service_ACCOUNTS:   "localhost:9001",
	discovery.Service_TODOGROUPS: "localhost:9002",
	discovery.Service_TODOS:      "localhost:9003",
}

func getGRPCConnection(service discovery.Service) (*grpc.ClientConn, error) {
	serverAddr := ServiceAddrMap[service]
	return grpc.Dial(serverAddr, grpc.WithInsecure())
}

func NewAuthServiceClient() (auth.AuthServiceClient, error) {
	clientConn, err := getGRPCConnection(discovery.Service_AUTH)
	if err != nil {
		return nil, err
	}
	return auth.NewAuthServiceClient(clientConn), nil
}

func NewAccountsServiceClient() (accounts.AccountsServiceClient, error) {
	clientConn, err := getGRPCConnection(discovery.Service_ACCOUNTS)
	if err != nil {
		return nil, err
	}
	return accounts.NewAccountsServiceClient(clientConn), nil
}

func NewTodoGroupsServiceClient() (todo_groups.TodoGroupsServiceClient, error) {
	clientConn, err := getGRPCConnection(discovery.Service_TODOGROUPS)
	if err != nil {
		return nil, err
	}
	return todo_groups.NewTodoGroupsServiceClient(clientConn), nil
}
