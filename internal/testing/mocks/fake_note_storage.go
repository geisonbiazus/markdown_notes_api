package mocks

import (
	"reflect"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type fakeNoteStorage struct {
	onSave func(markdownnotes.Note)
}

func (f *fakeNoteStorage) Save(n markdownnotes.Note) {
	f.onSave(n)
}

func (f *fakeNoteStorage) ShouldReceiveSaveWithNote(
	t *testing.T, note markdownnotes.Note,
) *bool {
	called := false
	f.onSave = func(n markdownnotes.Note) {
		if !reflect.DeepEqual(n, note) {
			t.Errorf("Wrong note argument. Expected: %v. Actual: %v", note, n)
		}
		called = true
	}
	return &called
}

func NewFakeNoteStorage() *fakeNoteStorage {
	return new(fakeNoteStorage)
}
