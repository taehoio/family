package grpc

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
	"github.com/taeho-io/family/services/base/aws"
	"github.com/taeho-io/family/services/base/aws/dynamodb"
	grpcService "github.com/taeho-io/family/services/base/grpc"
	"github.com/taeho-io/family/services/todo_groups/config"
	"github.com/taeho-io/family/services/todo_groups/grpc/handlers"
	"github.com/taeho-io/family/services/todo_groups/repos/todo_group_permits_repo"
	"github.com/taeho-io/family/services/todo_groups/repos/todo_groups_repo"
)

type IFace interface {
	grpcService.IFace

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
	todos.RegisterTodosServiceServer(srv, s)
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

func (s *Service) CreateTodoGroup(ctx context.Context, req *todos.CreateTodoGroupRequest) (*todos.CreateTodoGroupResponse, error) {
	return handlers.CreateTodoGroup(s.TodoGroupsTable())(ctx, req)
}

func (s *Service) GetTodoGroup(ctx context.Context, req *todos.GetTodoGroupRequest) (*todos.GetTodoGroupResponse, error) {
	return handlers.GetTodoGroup(s.TodoGroupsTable())(ctx, req)
}

func (s *Service) ListTodoGroups(ctx context.Context, req *todos.ListTodoGroupsRequest) (*todos.ListTodoGroupsResponse, error) {
	return handlers.ListTodoGroups(s.TodoGroupsTable())(ctx, req)
}
