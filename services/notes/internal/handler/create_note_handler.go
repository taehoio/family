package handler

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/notes"
	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/notes/internal/model"
	"github.com/taeho-io/family/services/notes/internal/repo"
)

type CreateNoteFunc func(ctx context.Context, req *notes.CreateNoteRequest) (*notes.CreateNoteResponse, error)

func CreateNote(notesRepo repo.NotesRepo) CreateNoteFunc {
	return func(ctx context.Context, req *notes.CreateNoteRequest) (*notes.CreateNoteResponse, error) {
		if err := validateCreateNoteInput(req); err != nil {
			return nil, err
		}

		accountID := base.GetAccountIDFromContext(ctx)
		if err := base.HasPermissionByAccountID(ctx, accountID); err != nil {
			return nil, err
		}

		logger := ctxlogrus.Extract(ctx).WithField("req", req)

		note := model.NewNoteFromProto(req.Note)
		note.NoteID = xid.New().String()
		note.CreateBy = accountID

		if err := notesRepo.Put(note); err != nil {
			logger.WithFields(logrus.Fields{
				"what": "notesRepo.Put",
				"note": note,
			}).Error(err)

			return nil, err
		}

		return &notes.CreateNoteResponse{
			Note: note.ToProto(),
		}, nil
	}
}

func validateCreateNoteInput(req *notes.CreateNoteRequest) error {
	if req.Note.Text == "" {
		return repo.InvalidTextError
	}

	return nil
}
