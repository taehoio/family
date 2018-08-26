package todos_service

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/taeho-io/family/idl/generated/go/pb/family/discovery"
	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
	"github.com/taeho-io/family/services/base/grpc/base_service"
	"github.com/taeho-io/family/services/base/grpc/dynamodb_service"
	"github.com/taeho-io/family/services/base/grpc/interceptors"
	"github.com/taeho-io/family/services/discovery/grpc/discovery_service"
	"github.com/taeho-io/family/services/todos/config"
	"github.com/taeho-io/family/services/todos/grpc/todos_service/handlers"
	"github.com/taeho-io/family/services/todos/repos/todos_repo"
)

type IFace interface {
	dynamodb_service.IFace
	todos.TodosServiceServer

	TodosTable() *todos_repo.Table
	TodoGroupsServiceClient() todo_groups.TodoGroupsServiceClient
}

type Service struct {
	dynamodb_service.IFace

	todosTable              todos_repo.IFace
	todoGroupsServiceClient todo_groups.TodoGroupsServiceClient
}

func New(cfg config.IFace) (*Service, error) {
	dynamodbSvc, err := dynamodb_service.New(cfg)
	if err != nil {
		return nil, err
	}

	todoGroupsServiceClient, err := discovery_service.NewTodoGroupsServiceClient()
	if err != nil {
		dynamodbSvc.Logger().WithFields(logrus.Fields{
			"what": "discovery_service.NewTodoGroupsServiceClient",
		}).Error(err)

		return nil, err
	}

	return &Service{
		IFace:                   dynamodbSvc,
		todosTable:              todos_repo.New(dynamodbSvc.Dynamodb(), cfg),
		todoGroupsServiceClient: todoGroupsServiceClient,
	}, nil
}

func NewMock() (*Service, error) {
	return New(config.NewMock())
}

func (s *Service) RegisterService(srv *grpc.Server) {
	todos.RegisterTodosServiceServer(srv, s)
}

func (s *Service) TodosTable() todos_repo.IFace {
	return s.todosTable
}

func (s *Service) TodoGroupsServiceClient() todo_groups.TodoGroupsServiceClient {
	return s.todoGroupsServiceClient
}

func (s *Service) CreateTodo(ctx context.Context, req *todos.CreateTodoRequest) (*todos.CreateTodoResponse, error) {
	return handlers.CreateTodo(s.TodosTable(), base_service.HasPermissionByAccountID)(ctx, req)
}

func (s *Service) GetTodo(ctx context.Context, req *todos.GetTodoRequest) (*todos.GetTodoResponse, error) {
	return handlers.GetTodo(
		s.TodosTable(),
		base_service.GetAccountIDFromContext,
		base_service.HasPermissionByAccountID,
		s.TodoGroupsServiceClient(),
	)(ctx, req)
}

func (s *Service) ListTodos(ctx context.Context, req *todos.ListTodosRequest) (*todos.ListTodosResponse, error) {
	return handlers.ListTodos(
		s.TodosTable(),
		base_service.GetAccountIDFromContext,
		base_service.HasPermissionByAccountID,
		s.TodoGroupsServiceClient(),
	)(ctx, req)
}

func (s *Service) UpdateTodo(ctx context.Context, req *todos.UpdateTodoRequest) (*todos.UpdateTodoResponse, error) {
	return handlers.UpdateTodo(
		s.TodosTable(),
		base_service.GetAccountIDFromContext,
		base_service.HasPermissionByAccountID,
		s.TodoGroupsServiceClient(),
	)(ctx, req)
}

func (s *Service) DeleteTodo(ctx context.Context, req *todos.DeleteTodoRequest) (*todos.DeleteTodoResponse, error) {
	return handlers.DeleteTodo(
		s.TodosTable(),
		base_service.GetAccountIDFromContext,
		base_service.HasPermissionByAccountID,
		s.TodoGroupsServiceClient(),
	)(ctx, req)
}

func Serve() error {
	addr := discovery_service.ServiceAddrMap[discovery.Service_TODOS]
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil
	}

	svr := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			interceptors.RequestIDUnaryServerInterceptor,
			interceptors.AuthUnaryServerInterceptor,
			interceptors.LogrusUnaryServerInterceptor,
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
