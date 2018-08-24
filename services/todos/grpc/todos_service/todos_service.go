package todos_service

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
	"github.com/taeho-io/family/services/base/aws/dynamodb"
	"github.com/taeho-io/family/services/base/grpc/dynamodb_service"
	"github.com/taeho-io/family/services/todos/config"
	"github.com/taeho-io/family/services/todos/grpc/todos_service/handlers"
	"github.com/taeho-io/family/services/todos/repos/todos_repo"
)

type IFace interface {
	dynamodb_service.IFace

	TodosTable() *todos_repo.Table
}

type Service struct {
	dynamodb_service.IFace

	cfg        config.IFace
	ddb        dynamodb.IFace
	todosTable *todos_repo.Table
}

func New(cfg config.IFace) (*Service, error) {
	dynamodbSvc, err := dynamodb_service.New(cfg)
	if err != nil {
		return nil, err
	}

	return &Service{
		IFace:      dynamodbSvc,
		todosTable: todos_repo.New(dynamodbSvc.Dynamodb(), cfg),
	}, nil
}

func NewMock() (*Service, error) {
	return New(config.NewMock())
}

func (s *Service) RegisterService(srv *grpc.Server) {
	todos.RegisterTodosServiceServer(srv, s)
}

func (s *Service) TodosTable() *todos_repo.Table {
	return s.todosTable
}

func (s *Service) CreateTodo(ctx context.Context, req *todos.CreateTodoRequest) (*todos.CreateTodoResponse, error) {
	return handlers.CreateTodo(s.TodosTable())(ctx, req)
}

func (s *Service) GetTodo(ctx context.Context, req *todos.GetTodoRequest) (*todos.GetTodoResponse, error) {
	return handlers.GetTodo(s.TodosTable())(ctx, req)
}

func (s *Service) ListTodos(ctx context.Context, req *todos.ListTodosRequest) (*todos.ListTodosResponse, error) {
	return handlers.ListTodos(s.TodosTable())(ctx, req)
}

func (s *Service) UpdateTodo(ctx context.Context, req *todos.UpdateTodoRequest) (*todos.UpdateTodoResponse, error) {
	return handlers.UpdateTodo(s.TodosTable())(ctx, req)
}

func (s *Service) DeleteTodo(ctx context.Context, req *todos.DeleteTodoRequest) (*todos.DeleteTodoResponse, error) {
	return handlers.DeleteTodo(s.TodosTable())(ctx, req)
}
