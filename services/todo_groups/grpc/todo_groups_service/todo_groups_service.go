package todo_groups_service

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/taeho-io/family/idl/generated/go/pb/family/discovery"
	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
	"github.com/taeho-io/family/services/base/aws"
	"github.com/taeho-io/family/services/base/aws/dynamodb"
	"github.com/taeho-io/family/services/base/grpc/base_service"
	"github.com/taeho-io/family/services/base/grpc/interceptors"
	"github.com/taeho-io/family/services/discovery/grpc/discovery_service"
	"github.com/taeho-io/family/services/todo_groups/config"
	"github.com/taeho-io/family/services/todo_groups/grpc/todo_groups_service/handlers"
	"github.com/taeho-io/family/services/todo_groups/repos/todo_group_permits_repo"
	"github.com/taeho-io/family/services/todo_groups/repos/todo_groups_repo"
)

type IFace interface {
	base_service.IFace

	Config() config.IFace
	DynamoDB() dynamodb.IFace
	TodoGroupsTable() *todo_groups_repo.Table
	TodoGroupPermitsTable() *todo_group_permits_repo.Table
}

type Service struct {
	IFace

	cfg                   config.IFace
	ddb                   dynamodb.IFace
	todoGroupsTable       *todo_groups_repo.Table
	todoGroupPermitsTable *todo_group_permits_repo.Table
}

func New(cfg config.IFace) (*Service, error) {
	a, err := aws.New()
	if err != nil {
		return nil, err
	}
	ddb := dynamodb.New(a)

	return &Service{
		cfg:                   cfg,
		ddb:                   ddb,
		todoGroupsTable:       todo_groups_repo.New(ddb, cfg),
		todoGroupPermitsTable: todo_group_permits_repo.New(ddb, cfg),
	}, nil
}

func NewMock() (*Service, error) {
	return New(config.NewMock())
}

func (s *Service) RegisterService(srv *grpc.Server) {
	todo_groups.RegisterTodoGroupsServiceServer(srv, s)
}

func (s *Service) Config() config.IFace {
	return s.cfg
}

func (s *Service) Dynamodb() dynamodb.IFace {
	return s.ddb
}

func (s *Service) TodoGroupsTable() *todo_groups_repo.Table {
	return s.todoGroupsTable
}

func (s *Service) TodoGroupPermitsTable() *todo_group_permits_repo.Table {
	return s.todoGroupPermitsTable
}

func (s *Service) CreateTodoGroup(
	ctx context.Context,
	req *todo_groups.CreateTodoGroupRequest,
) (*todo_groups.CreateTodoGroupResponse, error) {
	return handlers.CreateTodoGroup(s.TodoGroupsTable())(ctx, req)
}

func (s *Service) GetTodoGroup(
	ctx context.Context,
	req *todo_groups.GetTodoGroupRequest,
) (*todo_groups.GetTodoGroupResponse, error) {
	return handlers.GetTodoGroup(s.TodoGroupsTable())(ctx, req)
}

func (s *Service) ListTodoGroups(
	ctx context.Context,
	req *todo_groups.ListTodoGroupsRequest,
) (*todo_groups.ListTodoGroupsResponse, error) {
	return handlers.ListTodoGroups(s.TodoGroupsTable())(ctx, req)
}

func (s *Service) UpdateTodoGroup(
	ctx context.Context,
	req *todo_groups.UpdateTodoGroupRequest,
) (*todo_groups.UpdateTodoGroupResponse, error) {
	return handlers.UpdateTodoGroup(s.TodoGroupsTable())(ctx, req)
}

func (s *Service) DeleteTodoGroup(
	ctx context.Context,
	req *todo_groups.DeleteTodoGroupRequest,
) (*todo_groups.DeleteTodoGroupResponse, error) {
	return handlers.DeleteTodoGroup(s.TodoGroupsTable())(ctx, req)
}

func Serve() error {
	addr := discovery_service.ServiceAddrMap[discovery.Service_TODOGROUPS]
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	var whiteListPrefixes []string

	svr := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			interceptors.NewRequestIdIfNotExistsUnaryServerInterceptor,
			interceptors.ForwardRequestIdLogFieldUnaryServerInterceptor,
			interceptors.LogrusUnaryServerInterceptor,
			interceptors.AuthWhileListUnaryServerInterceptor(whiteListPrefixes),
			interceptors.AuthUnaryServerInterceptor,
		),
	)

	svc, err := New(config.New(config.NewSettings()))
	if err != nil {
		return err
	}

	svc.RegisterService(svr)
	reflection.Register(svr)
	return svr.Serve(lis)
}
