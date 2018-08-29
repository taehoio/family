package todos

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/taeho-io/family/idl/generated/go/pb/family/discovery"
	"github.com/taeho-io/family/idl/generated/go/pb/family/todogroups"
	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
	"github.com/taeho-io/family/services/base"
	discoveryService "github.com/taeho-io/family/services/discovery"
	"github.com/taeho-io/family/services/todos/internal/handler"
	"github.com/taeho-io/family/services/todos/internal/repo"
)

type Service interface {
	base.DynamodbService
	todos.TodosServiceServer

	TodosRepo() repo.TodosRepo
	TodoGroupsServiceClient() todogroups.TodoGroupsServiceClient
}

type defaultService struct {
	base.DynamodbService

	todosRepo               repo.TodosRepo
	todoGroupsServiceClient todogroups.TodoGroupsServiceClient
}

func New(cfg Config) (*defaultService, error) {
	dynamodbSvc, err := base.NewDynamodbService(cfg)
	if err != nil {
		return nil, err
	}

	todoGroupsServiceClient, err := discoveryService.NewTodoGroupsServiceClient()
	if err != nil {
		dynamodbSvc.Logger().WithFields(logrus.Fields{
			"what": "discovery_service.NewTodoGroupsServiceClient",
		}).Error(err)

		return nil, err
	}

	todosRepo := repo.NewTodosRepo(
		dynamodbSvc.Dynamodb(),
		repo.NewTodosRepoConfig(
			base.FullDynamodbTableName(cfg, cfg.Settings().DynamodbTodosTableName),
		),
	)

	return &defaultService{
		DynamodbService:         dynamodbSvc,
		todosRepo:               todosRepo,
		todoGroupsServiceClient: todoGroupsServiceClient,
	}, nil
}

func NewMock() (*defaultService, error) {
	return New(NewMockConfig())
}

func (s *defaultService) RegisterService(srv *grpc.Server) {
	todos.RegisterTodosServiceServer(srv, s)
}

func (s *defaultService) TodosRepo() repo.TodosRepo {
	return s.todosRepo
}

func (s *defaultService) TodoGroupsServiceClient() todogroups.TodoGroupsServiceClient {
	return s.todoGroupsServiceClient
}

func (s *defaultService) CreateTodo(ctx context.Context, req *todos.CreateTodoRequest) (*todos.CreateTodoResponse, error) {
	return handler.CreateTodo(s.TodosRepo(), base.HasPermissionByAccountID)(ctx, req)
}

func (s *defaultService) GetTodo(ctx context.Context, req *todos.GetTodoRequest) (*todos.GetTodoResponse, error) {
	return handler.GetTodo(
		s.TodosRepo(),
		base.GetAccountIDFromContext,
		base.HasPermissionByAccountID,
		s.TodoGroupsServiceClient(),
	)(ctx, req)
}

func (s *defaultService) ListTodos(ctx context.Context, req *todos.ListTodosRequest) (*todos.ListTodosResponse, error) {
	return handler.ListTodos(
		s.TodosRepo(),
		base.GetAccountIDFromContext,
		base.HasPermissionByAccountID,
		s.TodoGroupsServiceClient(),
	)(ctx, req)
}

func (s *defaultService) UpdateTodo(ctx context.Context, req *todos.UpdateTodoRequest) (*todos.UpdateTodoResponse, error) {
	return handler.UpdateTodo(
		s.TodosRepo(),
		base.GetAccountIDFromContext,
		base.HasPermissionByAccountID,
		s.TodoGroupsServiceClient(),
	)(ctx, req)
}

func (s *defaultService) DeleteTodo(ctx context.Context, req *todos.DeleteTodoRequest) (*todos.DeleteTodoResponse, error) {
	return handler.DeleteTodo(
		s.TodosRepo(),
		base.GetAccountIDFromContext,
		base.HasPermissionByAccountID,
		s.TodoGroupsServiceClient(),
	)(ctx, req)
}

func Serve() error {
	addr := discoveryService.ServiceAddrMap[discovery.Service_TODOS]
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil
	}

	svr := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			base.RequestIDUnaryServerInterceptor,
			base.AuthUnaryServerInterceptor,
			base.LogrusUnaryServerInterceptor,
		),
	)

	svc, err := New(NewConfig(NewSettings()))
	if err != nil {
		return err
	}

	svc.RegisterService(svr)
	reflection.Register(svr)
	return svr.Serve(lis)
}
