package doubles

import "github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"

type UpdatedNotePresenterSpy struct {
	PresentUpdatedNoteNoteArg markdownnotes.Note
	PresentErrorsErrsArg      []markdownnotes.ValidationError
	NotFoundCalled            bool
}

func NewUpdatedNotePresenterSpy() *UpdatedNotePresenterSpy {
	return new(UpdatedNotePresenterSpy)
}

func (s *UpdatedNotePresenterSpy) PresentUpdatedNote(note markdownnotes.Note) {
	s.PresentUpdatedNoteNoteArg = note
}

func (s *UpdatedNotePresenterSpy) PresentErrors(errs []markdownnotes.ValidationError) {
	s.PresentErrorsErrsArg = errs
}

func (s *UpdatedNotePresenterSpy) NotFound() {
	s.NotFoundCalled = true
}
