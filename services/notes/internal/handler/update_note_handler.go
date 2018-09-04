package handler

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/notes"
	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/notes/internal/repo"
)

type UpdateNoteFunc func(ctx context.Context, req *notes.UpdateNoteRequest) (*notes.UpdateNoteResponse, error)

func UpdateNote(notesRepo repo.NotesRepo) UpdateNoteFunc {
	return func(ctx context.Context, req *notes.UpdateNoteRequest) (*notes.UpdateNoteResponse, error) {
		accountID := base.GetAccountIDFromContext(ctx)
		if err := base.HasPermissionByAccountID(ctx, accountID); err != nil {
			return nil, err
		}

		if err := validateUpdateNoteInput(req); err != nil {
			return nil, err
		}

		logger := ctxlogrus.Extract(ctx).WithField("req", req)

		note, err := notesRepo.UpdateText(req.Note.NoteId, req.Note.TextType, req.Note.Text)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"what": "notesRepo.UpdateText",
			}).Error(err)

			return nil, err
		}

		return &notes.UpdateNoteResponse{
			Note: note.ToProto(),
		}, nil
	}
}

func validateUpdateNoteInput(req *notes.UpdateNoteRequest) error {
	if req.Note.NoteId == "" {
		return repo.InvalidNoteIDError
	}
	if req.Note.Text == "" {
		return repo.InvalidTextError
	}

	return nil
}
