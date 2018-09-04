package notes

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/taeho-io/family/idl/generated/go/pb/family/discovery"
	"github.com/taeho-io/family/idl/generated/go/pb/family/notes"
	"github.com/taeho-io/family/services/base"
	discoveryService "github.com/taeho-io/family/services/discovery"
	"github.com/taeho-io/family/services/notes/internal/handler"
	"github.com/taeho-io/family/services/notes/internal/repo"
)

type Service interface {
	base.DynamodbService
	notes.NotesServiceServer

	NotesRepo() repo.NotesRepo
}

type defaultService struct {
	base.DynamodbService

	notesRepo repo.NotesRepo
}

func New(cfg Config) (Service, error) {
	dynamodbSvc, err := base.NewDynamodbService(cfg)
	if err != nil {
		return nil, err
	}

	notesRepo := repo.NewNotesRepo(
		dynamodbSvc.Dynamodb(),
		repo.NewNotesRepoConfig(
			base.FullDynamodbTableName(cfg, cfg.Settings().DynamodbNotesTableName),
		),
	)

	return &defaultService{
		DynamodbService: dynamodbSvc,
		notesRepo:       notesRepo,
	}, nil
}

func NewMock() (Service, error) {
	return New(NewMockConfig())
}

func (s *defaultService) RegisterService(srv *grpc.Server) {
	notes.RegisterNotesServiceServer(srv, s)
}

func (s *defaultService) NotesRepo() repo.NotesRepo {
	return s.notesRepo
}

func (s *defaultService) CreateNote(ctx context.Context, req *notes.CreateNoteRequest) (*notes.CreateNoteResponse, error) {
	return handler.CreateNote(s.NotesRepo())(ctx, req)
}

func (s *defaultService) GetNote(ctx context.Context, req *notes.GetNoteRequest) (*notes.GetNoteResponse, error) {
	return handler.GetNote(s.NotesRepo())(ctx, req)
}

func (s *defaultService) ListNotes(ctx context.Context, req *notes.ListNotesRequest) (*notes.ListNotesResponse, error) {
	return handler.ListNotes(s.NotesRepo())(ctx, req)
}

func (s *defaultService) UpdateNote(ctx context.Context, req *notes.UpdateNoteRequest) (*notes.UpdateNoteResponse, error) {
	return handler.UpdateNote(s.NotesRepo())(ctx, req)
}

func (s *defaultService) DeleteNote(ctx context.Context, req *notes.DeleteNoteRequest) (*notes.DeleteNoteResponse, error) {
	return handler.DeleteNote(s.NotesRepo())(ctx, req)
}

func Serve() error {
	addr := discoveryService.ServiceAddrMap[discovery.Service_NOTES]
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
