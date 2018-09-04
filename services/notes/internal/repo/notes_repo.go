package repo

import (
	"fmt"
	"time"

	"github.com/taeho-io/family/idl/generated/go/pb/family/notes"

	"github.com/guregu/dynamo"

	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/notes/internal/model"
)

var (
	noteIDFieldKey     = "note_id"
	createdByFieldKey  = "created_by"
	createdByIndexName = "created_by-index"
	updatedAtFieldKey  = "updated_at"
	textTypeFieldKey   = "text_type"
	textFieldKey       = "text"
)

var (
	InvalidNoteError     = fmt.Errorf("invalid note")
	InvalidNoteIDError   = fmt.Errorf("invliad note_id")
	InvalidCreateByError = fmt.Errorf("invalid create_by")
	InvalidTextError     = fmt.Errorf("invalid text")
)

type NotesRepo interface {
	Put(note *model.Note) error
	GetByID(noteID string) (*model.Note, error)
	ListByCreatedBy(createdBy string) ([]*model.Note, error)
	UpdateText(noteID string, textType notes.TextType, text string) (*model.Note, error)
	DeleteByID(noteID string) error
}

type dynamodbNotesRepo struct {
	NotesRepo
	base.DynamodbRepo

	notesTable *dynamo.Table
}

func NewNotesRepo(ddb base.Dynamodb, cfg NotesRepoConfig) NotesRepo {
	notesTable := ddb.DB().Table(cfg.FullTableName())

	return &dynamodbNotesRepo{
		notesTable: &notesTable,
	}
}

func NewMockNotesRepo() NotesRepo {
	ddb := base.NewMockDynamodb()
	cfg := NewMockNotesRepoConfig()

	return NewNotesRepo(ddb, cfg)
}

func validateNoteInput(note *model.Note) error {
	if note == nil {
		return InvalidNoteError
	}
	if note.NoteID == "" {
		return InvalidNoteIDError
	}
	if note.CreateBy == "" {
		return InvalidCreateByError
	}
	if note.Text == "" {
		return InvalidTextError
	}

	return nil
}

func (n *dynamodbNotesRepo) Table() *dynamo.Table {
	return n.notesTable
}

func (n *dynamodbNotesRepo) Put(note *model.Note) error {
	if err := validateNoteInput(note); err != nil {
		return err
	}

	now := time.Now()
	note.CreatedAt = now
	note.UpdatedAt = now

	return n.Table().Put(note).Run()
}

func (n *dynamodbNotesRepo) GetByID(noteID string) (*model.Note, error) {
	var note model.Note

	if err := n.Table().Get(noteIDFieldKey, noteID).One(&note); err != nil {
		return nil, err
	}

	return &note, nil
}

func (n *dynamodbNotesRepo) ListByCreatedBy(createdBy string) ([]*model.Note, error) {
	var noteList []*model.Note

	err := n.Table().
		Get(createdByFieldKey, createdBy).
		Index(createdByIndexName).
		Order(false).
		All(&noteList)
	if err != nil {
		return nil, err
	}

	return noteList, nil
}

func (n *dynamodbNotesRepo) UpdateText(noteID string, textType notes.TextType, text string) (*model.Note, error) {
	var note model.Note

	err := n.Table().
		Update(noteIDFieldKey, noteID).
		If(fmt.Sprintf("%s = ?", noteIDFieldKey), noteID).
		Set(textTypeFieldKey, textType).
		Set(textFieldKey, text).
		Set(updatedAtFieldKey, time.Now()).
		Value(&note)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (n *dynamodbNotesRepo) DeleteByID(noteID string) error {
	return n.Table().
		Delete(noteIDFieldKey, noteID).
		If(fmt.Sprintf("%s = ?", noteIDFieldKey), noteID).
		Run()
}
