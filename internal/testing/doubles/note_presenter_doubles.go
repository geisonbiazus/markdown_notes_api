package doubles

import "github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"

type NotePresenterSpy struct {
	PresentNoteNoteArg  markdownnotes.Note
	PresentErrorErrsArg []markdownnotes.ValidationError
	NotFoundCalled      bool
}

func NewNotePresenterSpy() *NotePresenterSpy {
	return new(NotePresenterSpy)
}

func (s *NotePresenterSpy) PresentNote(note markdownnotes.Note) {
	s.PresentNoteNoteArg = note
}

func (s *NotePresenterSpy) NotFound() {
	s.NotFoundCalled = true
}
