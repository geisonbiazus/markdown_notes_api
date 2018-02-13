package doubles

import (
	"errors"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type NoteStorageSpy struct {
	SaveNoteResult markdownnotes.Note
	saveNoteArg    markdownnotes.Note
}

func (s *NoteStorageSpy) Save(n markdownnotes.Note) (markdownnotes.Note, error) {
	s.saveNoteArg = n
	return s.SaveNoteResult, nil
}

func (s *NoteStorageSpy) SaveNoteArg() markdownnotes.Note {
	return s.saveNoteArg
}

func NewNoteStorageSpy() *NoteStorageSpy {
	return &NoteStorageSpy{}
}

type ErrorNoteStorageStub struct {
	Error error
}

func (s *ErrorNoteStorageStub) Save(n markdownnotes.Note) (markdownnotes.Note, error) {
	return markdownnotes.Note{}, s.Error
}

func NewErrorNoteSotorageStub() *ErrorNoteStorageStub {
	return &ErrorNoteStorageStub{errors.New("Some Error")}
}
