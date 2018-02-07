package mocks

import (
	"reflect"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type fakeNoteStorage struct {
	onSave func(markdownnotes.Note) error
}

func (f *fakeNoteStorage) Save(n markdownnotes.Note) error {
	return f.onSave(n)
}

func (f *fakeNoteStorage) ShouldReceiveSaveWithNoteAndReturn(
	t *testing.T, note markdownnotes.Note, result error,
) *bool {
	called := false
	f.onSave = func(n markdownnotes.Note) error {
		if !reflect.DeepEqual(n, note) {
			t.Errorf("Wrong note argument. Expected: %v. Actual: %v", note, n)
		}
		called = true
		return result
	}
	return &called
}

func NewFakeNoteStorage() *fakeNoteStorage {
	return new(fakeNoteStorage)
}
