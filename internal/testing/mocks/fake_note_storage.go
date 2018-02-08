package mocks

import (
	"reflect"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type FakeNoteStorage struct {
	onSave func(markdownnotes.Note) (markdownnotes.Note, error)
}

func (f *FakeNoteStorage) Save(n markdownnotes.Note) (markdownnotes.Note, error) {
	return f.onSave(n)
}

func (f *FakeNoteStorage) ShouldReceiveSaveWithNoteAndReturn(
	t *testing.T, noteArg markdownnotes.Note,
	noteResult markdownnotes.Note, errResult error,
) *bool {
	called := false
	f.onSave = func(n markdownnotes.Note) (markdownnotes.Note, error) {
		if !reflect.DeepEqual(n, noteArg) {
			t.Errorf("Wrong note argument. Expected: %v. Actual: %v", noteArg, n)
		}
		called = true
		return noteResult, errResult
	}
	return &called
}

func (f *FakeNoteStorage) ShouldNotReceiveSave(t *testing.T) {
	f.onSave = func(n markdownnotes.Note) (markdownnotes.Note, error) {
		t.Error("Save was called when it shouldn't")
		return n, nil
	}
}

func NewFakeNoteStorage() *FakeNoteStorage {
	return &FakeNoteStorage{
		onSave: func(markdownnotes.Note) (markdownnotes.Note, error) {
			return markdownnotes.Note{}, nil
		},
	}
}
