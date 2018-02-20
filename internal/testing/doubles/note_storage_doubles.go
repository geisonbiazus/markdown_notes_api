package doubles

import (
	"errors"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type NoteStorageSpy struct {
	SaveNoteResult markdownnotes.Note
	SaveNoteArg    markdownnotes.Note
	FindAllResult  []markdownnotes.Note
}

func NewNoteStorageSpy() *NoteStorageSpy {
	return &NoteStorageSpy{}
}

func (s *NoteStorageSpy) Save(n markdownnotes.Note) (markdownnotes.Note, error) {
	s.SaveNoteArg = n
	return s.SaveNoteResult, nil
}

func (s *NoteStorageSpy) FindAll() ([]markdownnotes.Note, error) {
	return s.FindAllResult, nil
}

type ErrorNoteStorageStub struct {
	Error error
}

func NewErrorNoteSotorageStub() *ErrorNoteStorageStub {
	return &ErrorNoteStorageStub{errors.New("Some Error")}
}

func (s *ErrorNoteStorageStub) Save(n markdownnotes.Note) (markdownnotes.Note, error) {
	return markdownnotes.Note{}, s.Error
}

func (s *ErrorNoteStorageStub) FindAll() ([]markdownnotes.Note, error) {
	return []markdownnotes.Note{}, s.Error
}
