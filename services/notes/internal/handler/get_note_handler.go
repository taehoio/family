package handler

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/notes"
	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/notes/internal/repo"
)

type GetNoteFunc func(ctx context.Context, req *notes.GetNoteRequest) (*notes.GetNoteResponse, error)

func GetNote(notesRepo repo.NotesRepo) GetNoteFunc {
	return func(ctx context.Context, req *notes.GetNoteRequest) (*notes.GetNoteResponse, error) {
		accountID := base.GetAccountIDFromContext(ctx)
		if err := base.HasPermissionByAccountID(ctx, accountID); err != nil {
			return nil, err
		}

		logger := ctxlogrus.Extract(ctx).WithField("req", req)

		note, err := notesRepo.GetByID(req.NoteId)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"what": "notesRepo.GetByID",
			}).Error(err)

			return nil, err
		}

		return &notes.GetNoteResponse{
			Note: note.ToProto(),
		}, nil
	}
}
