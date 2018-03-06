package doubles

import "github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"

type CreatedNotePresenterSpy struct {
	PresentCreatedNoteNoteArg markdownnotes.Note
	PresentErrorsErrsArg      []markdownnotes.ValidationError
	NotFoundCalled            bool
}

func NewCreatedNotePresenterSpy() *CreatedNotePresenterSpy {
	return new(CreatedNotePresenterSpy)
}

func (s *CreatedNotePresenterSpy) PresentCreatedNote(note markdownnotes.Note) {
	s.PresentCreatedNoteNoteArg = note
}

func (s *CreatedNotePresenterSpy) PresentErrors(errs []markdownnotes.ValidationError) {
	s.PresentErrorsErrsArg = errs
}
