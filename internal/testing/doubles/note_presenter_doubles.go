package doubles

import "github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"

type NotePresenterSpy struct {
	PresentNoteNoteArg       markdownnotes.Note
	PresentErrorsErrsArg     []markdownnotes.ValidationError
	PresentNotesCalled       bool
	PresentNotesNotesArg     []markdownnotes.Note
	NotFoundCalled           bool
	ServiceUnavailableCalled bool
}

func NewNotePresenterSpy() *NotePresenterSpy {
	return new(NotePresenterSpy)
}

func (s *NotePresenterSpy) PresentNote(note markdownnotes.Note) {
	s.PresentNoteNoteArg = note
}

func (s *NotePresenterSpy) PresentNotes(n []markdownnotes.Note) {
	s.PresentNotesCalled = true
	s.PresentNotesNotesArg = n
}

func (s *NotePresenterSpy) PresentErrors(errs []markdownnotes.ValidationError) {
	s.PresentErrorsErrsArg = errs
}

func (s *NotePresenterSpy) NotFound() {
	s.NotFoundCalled = true
}

func (s *NotePresenterSpy) ServiceUnavailable() {
	s.ServiceUnavailableCalled = true
}
