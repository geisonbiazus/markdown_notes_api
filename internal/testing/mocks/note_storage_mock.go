package mocks

import (
	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
	"github.com/stretchr/testify/mock"
)

type NoteStorageMock struct {
	mock.Mock
}

func (f *NoteStorageMock) Save(n markdownnotes.Note) (markdownnotes.Note, error) {
	args := f.Called(n)
	return args.Get(0).(markdownnotes.Note), args.Error(1)
}

func NewNoteStorageMock() *NoteStorageMock {
	return new(NoteStorageMock)
}
