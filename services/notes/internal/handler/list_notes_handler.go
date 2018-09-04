package handler

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/notes"
	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/notes/internal/repo"
)

type ListNotesFunc func(ctx context.Context, req *notes.ListNotesRequest) (*notes.ListNotesResponse, error)

func ListNotes(notesRepo repo.NotesRepo) ListNotesFunc {
	return func(ctx context.Context, req *notes.ListNotesRequest) (*notes.ListNotesResponse, error) {
		accountID := base.GetAccountIDFromContext(ctx)
		if err := base.HasPermissionByAccountID(ctx, accountID); err != nil {
			return nil, err
		}

		logger := ctxlogrus.Extract(ctx).WithField("req", req)

		noteList, err := notesRepo.ListByCreatedBy(accountID)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"what":      "notesRepo.ListByCreatedBy",
				"accountID": accountID,
			}).Error(err)

			return nil, err
		}

		var noteProtos []*notes.Note
		for _, note := range noteList {
			noteProtos = append(noteProtos, note.ToProto())
		}

		return &notes.ListNotesResponse{
			Notes: noteProtos,
		}, nil
	}
}
