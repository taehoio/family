package model

import (
	"time"

	"github.com/taeho-io/family/idl/generated/go/pb/family/notes"
)

type Note struct {
	NoteID    string          `dynamo:"note_id,hash"`
	CreateBy  string          `dynamo:"created_by" index:"created_by-index,hash"`
	Text      string          `dynamo:"text"`
	CreatedAt time.Time       `dynamo:"created_at"`
	UpdatedAt time.Time       `dynamo:"updated_at" index:"created_by-index,range"`
	ShareType notes.ShareType `dynamo:"share_type"`
	TextType  notes.TextType  `dynamo:"text_type"`
}

func NewNoteFromProto(noteProto *notes.Note) *Note {
	return &Note{
		NoteID:    noteProto.NoteId,
		CreateBy:  noteProto.CreatedBy,
		Text:      noteProto.Text,
		CreatedAt: time.Unix(noteProto.CreatedAt, 0),
		UpdatedAt: time.Unix(noteProto.UpdatedAt, 0),
		ShareType: noteProto.ShareType,
		TextType:  noteProto.TextType,
	}
}

func (n *Note) ToProto() *notes.Note {
	return &notes.Note{
		NoteId:    n.NoteID,
		CreatedBy: n.CreateBy,
		Text:      n.Text,
		CreatedAt: n.CreatedAt.Unix(),
		UpdatedAt: n.UpdatedAt.Unix(),
		ShareType: n.ShareType,
		TextType:  n.TextType,
	}
}
