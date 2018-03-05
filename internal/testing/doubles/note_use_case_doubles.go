package doubles

import (
	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type NoteUseCaseSpy struct {
	CreateNoteTitleArg     string
	CreateNoteContentArg   string
	CreateNotePresenterArg markdownnotes.NotePresenter
	CreateNoteErrorResult  error
}

func (s *NoteUseCaseSpy) CreateNote(title, content string, presenter markdownnotes.NotePresenter) error {
	s.CreateNoteTitleArg = title
	s.CreateNoteContentArg = content
	s.CreateNotePresenterArg = presenter
	return s.CreateNoteErrorResult
}

func NewNoteUseCaseSpy() *NoteUseCaseSpy {
	return &NoteUseCaseSpy{}
}
