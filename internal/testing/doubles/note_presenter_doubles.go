package doubles

import "github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"

type NotePresenterSpy struct {
	PresentNoteNoteArg  markdownnotes.Note
	PresentErrorErrsArg []markdownnotes.ValidationError
}

func (s *NotePresenterSpy) PresentNote(note markdownnotes.Note) {
	s.PresentNoteNoteArg = note
}

func (s *NotePresenterSpy) PresentError(errs []markdownnotes.ValidationError) {
	s.PresentErrorErrsArg = errs
}

func NewNotePresenterSpy() *NotePresenterSpy {
	return new(NotePresenterSpy)
}
