package handler

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/notes"
	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/notes/internal/repo"
)

type DeleteNoteFunc func(ctx context.Context, req *notes.DeleteNoteRequest) (*notes.DeleteNoteResponse, error)

func DeleteNote(notesRepo repo.NotesRepo) DeleteNoteFunc {
	return func(ctx context.Context, req *notes.DeleteNoteRequest) (*notes.DeleteNoteResponse, error) {
		accountID := base.GetAccountIDFromContext(ctx)
		if err := base.HasPermissionByAccountID(ctx, accountID); err != nil {
			return nil, err
		}

		logger := ctxlogrus.Extract(ctx).WithField("req", req)

		if err := notesRepo.DeleteByID(req.NoteId); err != nil {
			logger.WithFields(logrus.Fields{
				"what": "notesRepo.DeleteByID",
			}).Error(err)
		}

		return &notes.DeleteNoteResponse{}, nil
	}
}
