package doubles

import "github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes/notes"

type CreateNoteUseCaseSpy struct {
	RunTitleArg     string
	RunContentArg   string
	RunPresenterArg notes.NotePresenter
}

func (s *CreateNoteUseCaseSpy) Run(title, content string, presenter notes.NotePresenter) error {
	s.RunTitleArg = title
	s.RunContentArg = content
	s.RunPresenterArg = presenter
	return nil
}

func NewCreateNoteUseCaseSpy() *CreateNoteUseCaseSpy {
	return &CreateNoteUseCaseSpy{}
}
