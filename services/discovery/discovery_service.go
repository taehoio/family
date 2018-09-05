package discovery

import (
	"google.golang.org/grpc"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/idl/generated/go/pb/family/discovery"
	"github.com/taeho-io/family/idl/generated/go/pb/family/notes"
	"github.com/taeho-io/family/idl/generated/go/pb/family/todogroups"
	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
)

var ServiceAddrMap = map[discovery.Service]string{
	discovery.Service_AUTH:       "localhost:9000",
	discovery.Service_ACCOUNTS:   "localhost:9001",
	discovery.Service_TODOGROUPS: "localhost:9002",
	discovery.Service_TODOS:      "localhost:9003",
	discovery.Service_NOTES:      "localhost:9004",
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

func NewTodoGroupsServiceClient() (todogroups.TodoGroupsServiceClient, error) {
	clientConn, err := getGRPCConnection(discovery.Service_TODOGROUPS)
	if err != nil {
		return nil, err
	}
	return todogroups.NewTodoGroupsServiceClient(clientConn), nil
}

func NewTodosServiceClient() (todos.TodosServiceClient, error) {
	clientConn, err := getGRPCConnection(discovery.Service_TODOS)
	if err != nil {
		return nil, err
	}
	return todos.NewTodosServiceClient(clientConn), nil
}

func NewNotesServiceClient() (notes.NotesServiceClient, error) {
	clientConn, err := getGRPCConnection(discovery.Service_NOTES)
	if err != nil {
		return nil, err
	}
	return notes.NewNotesServiceClient(clientConn), nil
}
