package doubles

import "github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"

type ListNotesUseCaseSpy struct {
	RunPresenterArg markdownnotes.NoteListPresenter
	RunErrorResult  error
}

func NewListNotesUseCaseSpy() *ListNotesUseCaseSpy {
	return &ListNotesUseCaseSpy{}
}

func (s *ListNotesUseCaseSpy) Run(p markdownnotes.NoteListPresenter) error {
	s.RunPresenterArg = p
	return s.RunErrorResult
}
