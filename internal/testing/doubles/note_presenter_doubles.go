package doubles

import "github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"

type NotePresenterSpy struct {
	PresentCreatedNoteCalled  bool
	PresentCreatedNoteNoteArg markdownnotes.Note

	PresentUpdatedNoteCalled  bool
	PresentUpdatedNoteNoteArg markdownnotes.Note

	PresentNoteCalled  bool
	PresentNoteNoteArg markdownnotes.Note

	PresentValidationErrorsCalled  bool
	PresentValidationErrorsErrsArg []markdownnotes.ValidationError

	PresentNoteListCalled   bool
	PresentNoteListNotesArg []markdownnotes.Note

	NotFoundCalled bool

	PresentErrorCalled   bool
	PresentErrorErrorArg error
}

func NewNotePresenterSpy() *NotePresenterSpy {
	return new(NotePresenterSpy)
}

func (s *NotePresenterSpy) PresentCreatedNote(note markdownnotes.Note) {
	s.PresentCreatedNoteCalled = true
	s.PresentCreatedNoteNoteArg = note
}

func (s *NotePresenterSpy) PresentUpdatedNote(note markdownnotes.Note) {
	s.PresentCreatedNoteCalled = true
	s.PresentUpdatedNoteNoteArg = note
}

func (s *NotePresenterSpy) PresentNote(note markdownnotes.Note) {
	s.PresentNoteCalled = true
	s.PresentNoteNoteArg = note
}

func (s *NotePresenterSpy) PresentNoteList(n []markdownnotes.Note) {
	s.PresentNoteListCalled = true
	s.PresentNoteListNotesArg = n
}

func (s *NotePresenterSpy) PresentValidationErrors(errs []markdownnotes.ValidationError) {
	s.PresentValidationErrorsCalled = true
	s.PresentValidationErrorsErrsArg = errs
}

func (s *NotePresenterSpy) PresentNotFound() {
	s.NotFoundCalled = true
}

func (s *NotePresenterSpy) PresentError(e error) {
	s.PresentErrorCalled = true
	s.PresentErrorErrorArg = e
}
