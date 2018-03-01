package doubles

import "github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"

type ShowNoteUseCaseSpy struct {
	RunNoteIDArg    int
	RunPresenterArg markdownnotes.NotePresenter
	RunErrorResult  error
}

func NewShowNoteUseCaseSpy() *ShowNoteUseCaseSpy {
	return &ShowNoteUseCaseSpy{}
}

func (s *ShowNoteUseCaseSpy) Run(noteID int, p markdownnotes.NotePresenter) error {
	s.RunNoteIDArg = noteID
	s.RunPresenterArg = p
	return s.RunErrorResult
}
