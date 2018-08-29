package todogroups

import (
	"net"

	"github.com/taeho-io/family/services/todogroups/internal/repo"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todogroups"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/taeho-io/family/idl/generated/go/pb/family/discovery"
	"github.com/taeho-io/family/services/base"
	discoveryService "github.com/taeho-io/family/services/discovery"
	"github.com/taeho-io/family/services/todogroups/internal/handler"
)

type Service interface {
	base.DynamodbService
	todogroups.TodoGroupsServiceServer

	TodoGroupsRepo() repo.GroupsRepo
	TodoGroupPermitsRepo() repo.PermitsRepo
}

type DefaultService struct {
	base.DynamodbService

	todoGroupsRepo       repo.GroupsRepo
	todoGroupPermitsRepo repo.PermitsRepo
}

func New(cfg Config) (Service, error) {
	dynamodbSvc, err := base.NewDynamodbService(cfg)
	if err != nil {
		return nil, err
	}

	todoGroupsRepo := repo.NewTodosRepo(
		dynamodbSvc.Dynamodb(),
		repo.NewTodoGroupsRepoConfig(
			base.FullDynamodbTableName(cfg, cfg.Settings().DynamodbTodoGroupsTableName),
		),
	)

	todoGroupPermitsRepo := repo.NewPermitsRepo(
		dynamodbSvc.Dynamodb(),
		repo.NewTodoGroupPermitsRepoConfig(
			base.FullDynamodbTableName(cfg, cfg.Settings().DynamodbTodoGroupPermitsTableName),
		),
	)

	return &DefaultService{
		DynamodbService:      dynamodbSvc,
		todoGroupsRepo:       todoGroupsRepo,
		todoGroupPermitsRepo: todoGroupPermitsRepo,
	}, nil
}

func NewMock() (Service, error) {
	return New(NewMockConfig())
}

func (s *DefaultService) RegisterService(srv *grpc.Server) {
	todogroups.RegisterTodoGroupsServiceServer(srv, s)
}

func (s *DefaultService) TodoGroupsRepo() repo.GroupsRepo {
	return s.todoGroupsRepo
}

func (s *DefaultService) TodoGroupPermitsRepo() repo.PermitsRepo {
	return s.todoGroupPermitsRepo
}

func (s *DefaultService) CreateTodoGroup(
	ctx context.Context,
	req *todogroups.CreateTodoGroupRequest,
) (
	*todogroups.CreateTodoGroupResponse,
	error,
) {
	return handler.CreateTodoGroup(
		s.TodoGroupsRepo(),
		s.TodoGroupPermitsRepo(),
		base.HasPermissionByAccountID,
	)(ctx, req)
}

func (s *DefaultService) GetTodoGroup(
	ctx context.Context,
	req *todogroups.GetTodoGroupRequest,
) (
	*todogroups.GetTodoGroupResponse,
	error,
) {
	return handler.GetTodoGroup(
		s.TodoGroupsRepo(),
		s.TodoGroupPermitsRepo(),
		base.GetAccountIDFromContext,
		base.HasPermissionByAccountID,
	)(ctx, req)
}

func (s *DefaultService) ListTodoGroups(
	ctx context.Context,
	req *todogroups.ListTodoGroupsRequest,
) (
	*todogroups.ListTodoGroupsResponse,
	error,
) {
	return handler.ListTodoGroups(
		s.TodoGroupsRepo(),
		s.TodoGroupPermitsRepo(),
		base.HasPermissionByAccountID,
	)(ctx, req)
}

func (s *DefaultService) UpdateTodoGroup(
	ctx context.Context,
	req *todogroups.UpdateTodoGroupRequest,
) (
	*todogroups.UpdateTodoGroupResponse,
	error,
) {
	return handler.UpdateTodoGroup(
		s.TodoGroupsRepo(),
		base.HasPermissionByAccountID,
	)(ctx, req)
}

func (s *DefaultService) DeleteTodoGroup(
	ctx context.Context,
	req *todogroups.DeleteTodoGroupRequest,
) (
	*todogroups.DeleteTodoGroupResponse,
	error,
) {
	return handler.DeleteTodoGroup(
		s.TodoGroupsRepo(),
		s.TodoGroupPermitsRepo(),
		base.HasPermissionByAccountID,
	)(ctx, req)
}

func Serve() error {
	addr := discoveryService.ServiceAddrMap[discovery.Service_TODOGROUPS]
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
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
