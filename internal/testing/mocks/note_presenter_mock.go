package mocks

import (
	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
	"github.com/stretchr/testify/mock"
)

type NotePresenterMock struct {
	mock.Mock
}

func (f *NotePresenterMock) PresentNote(note markdownnotes.Note) {
	f.Called(note)
}

func (f *NotePresenterMock) PresentError(errs []markdownnotes.ValidationError) {
	f.Called(errs)
}

func NewNotePresenterMock() *NotePresenterMock {
	return new(NotePresenterMock)
}
