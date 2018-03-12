package doubles

import (
	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type NoteUseCaseSpy struct {
	CreateNoteTitleArg     string
	CreateNoteContentArg   string
	CreateNotePresenterArg markdownnotes.NotePresenter
	CreateNoteErrorResult  error

	ListNotesPresenterArg markdownnotes.NotePresenter
	ListNotesErrorResult  error

	ShowNoteNoteIDArg    int
	ShowNotePresenterArg markdownnotes.NotePresenter
	ShowNoteErrorResult  error
}

func NewNoteUseCaseSpy() *NoteUseCaseSpy {
	return &NoteUseCaseSpy{}
}

func (s *NoteUseCaseSpy) CreateNote(title, content string, presenter markdownnotes.NotePresenter) error {
	s.CreateNoteTitleArg = title
	s.CreateNoteContentArg = content
	s.CreateNotePresenterArg = presenter
	return s.CreateNoteErrorResult
}

func (s *NoteUseCaseSpy) ListNotes(p markdownnotes.NotePresenter) error {
	s.ListNotesPresenterArg = p
	return s.ListNotesErrorResult
}

func (s *NoteUseCaseSpy) ShowNote(noteID int, p markdownnotes.NotePresenter) error {
	s.ShowNoteNoteIDArg = noteID
	s.ShowNotePresenterArg = p
	return s.ShowNoteErrorResult
}

func (s *NoteUseCaseSpy) UpdateNote(id int, title, content string, p markdownnotes.NotePresenter) error {
	return nil
}
