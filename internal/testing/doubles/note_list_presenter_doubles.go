package doubles

import "github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"

type NoteListPresenterSpy struct {
	PresentNotesCalled   bool
	PresentNotesNotesArg []markdownnotes.Note
}

func NewNoteListPresenterSpy() *NoteListPresenterSpy {
	return &NoteListPresenterSpy{}
}

func (s *NoteListPresenterSpy) PresentNotes(n []markdownnotes.Note) {
	s.PresentNotesCalled = true
	s.PresentNotesNotesArg = n
}
