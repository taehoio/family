package repo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/taeho-io/family/idl/generated/go/pb/family/notes"
	"github.com/taeho-io/family/services/notes/internal/model"
)

var (
	notesRepo NotesRepo

	testNoteID          = "test_node_id"
	testAnotherNoteID   = "test_another_node_id"
	testCreatedBy       = "test_created_by"
	testText            = "test_text"
	testUpdatedText     = "test_updated_text"
	testUpdatedTextType = notes.TextType_MARKDOWN
)

func TestMain(m *testing.M) {
	notesRepo = NewMockNotesRepo()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestValidateNoteInput(t *testing.T) {
	note := &model.Note{
		NoteID:    testNoteID,
		CreateBy:  testCreatedBy,
		Text:      testText,
		ShareType: notes.ShareType_PRIVATE,
		TextType:  notes.TextType_TEXT,
	}
	err := validateNoteInput(note)
	assert.Nil(t, err)
}

func TestPut(t *testing.T) {
	note := &model.Note{
		NoteID:    testNoteID,
		CreateBy:  testCreatedBy,
		Text:      testText,
		ShareType: notes.ShareType_PRIVATE,
		TextType:  notes.TextType_TEXT,
	}
	err := notesRepo.Put(note)
	assert.Nil(t, err)
}

func TestGet(t *testing.T) {
	note, err := notesRepo.GetByID(testNoteID)
	assert.NotNil(t, note)
	assert.Nil(t, err)
	assert.Equal(t, testNoteID, note.NoteID)
	assert.Equal(t, testText, note.Text)
}

func TestPutAnother(t *testing.T) {
	note := &model.Note{
		NoteID:    testAnotherNoteID,
		CreateBy:  testCreatedBy,
		Text:      testText,
		ShareType: notes.ShareType_PRIVATE,
		TextType:  notes.TextType_TEXT,
	}
	err := notesRepo.Put(note)
	assert.Nil(t, err)
}

func TestList(t *testing.T) {
	note := &model.Note{
		NoteID:    testAnotherNoteID,
		CreateBy:  testCreatedBy,
		Text:      testText,
		ShareType: notes.ShareType_PRIVATE,
		TextType:  notes.TextType_TEXT,
	}
	err := notesRepo.Put(note)
	assert.Nil(t, err)

	noteList, err := notesRepo.ListByCreatedBy(testCreatedBy)
	assert.NotNil(t, noteList)
	assert.Nil(t, err)
	assert.Len(t, noteList, 2)
	assert.Equal(t, testNoteID, noteList[1].NoteID)
	assert.Equal(t, testAnotherNoteID, noteList[0].NoteID)

}

func TestUpdateText(t *testing.T) {
	note, err := notesRepo.UpdateText(testNoteID, testUpdatedTextType, testUpdatedText)
	assert.NotNil(t, note)
	assert.Nil(t, err)
	assert.Equal(t, testUpdatedTextType, note.TextType)
	assert.Equal(t, testUpdatedText, note.Text)
	assert.NotEqual(t, note.CreatedAt, note.UpdatedAt)
}

func TestDelete(t *testing.T) {
	err := notesRepo.DeleteByID(testNoteID)
	assert.Nil(t, err)
}

func TestDeleteAnother(t *testing.T) {
	err := notesRepo.DeleteByID(testAnotherNoteID)
	assert.Nil(t, err)
}
