package doubles

import (
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type NotePresenterFactorySpy struct {
	CreateCalled            bool
	CreateResponseWriterArg http.ResponseWriter
	ReturnedNotePresenter   *NotePresenterSpy
}

func NewNotePresenterFactorySpy() *NotePresenterFactorySpy {
	return &NotePresenterFactorySpy{ReturnedNotePresenter: NewNotePresenterSpy()}
}

func (s *NotePresenterFactorySpy) Create(w http.ResponseWriter) markdownnotes.NotePresenter {
	s.CreateCalled = true
	s.CreateResponseWriterArg = w
	return s.ReturnedNotePresenter
}
