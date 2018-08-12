package discovery

import (
	"google.golang.org/grpc"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/idl/generated/go/pb/family/discovery"
)

var serverAddrMap = map[discovery.Service]string{
	discovery.Service_AUTH:    "localhost:9001",
	discovery.Service_ACCOUNT: "localhost:9001",
}

func getGRPCConnection(service discovery.Service) (*grpc.ClientConn, error) {
	serverAddr := serverAddrMap[service]
	return grpc.Dial(serverAddr, grpc.WithInsecure())
}

func NewAuthServiceClient() (auth.AuthServiceClient, error) {
	clientConn, err := getGRPCConnection(discovery.Service_AUTH)
	if err != nil {
		return nil, err
	}
	return auth.NewAuthServiceClient(clientConn), nil
}

func NewAccountServiceClient() (accounts.AccountsServiceClient, error) {
	clientConn, err := getGRPCConnection(discovery.Service_ACCOUNT)
	if err != nil {
		return nil, err
	}
	return accounts.NewAccountsServiceClient(clientConn), nil
}
