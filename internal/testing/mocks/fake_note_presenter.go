package mocks

import (
	"reflect"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type FakeNotePresenter struct {
	onPresentNote func(note markdownnotes.Note)
}

func (f *FakeNotePresenter) PresentNote(note markdownnotes.Note) {
	f.onPresentNote(note)
}

func (f *FakeNotePresenter) ShouldReceivePresentNoteWithNote(
	t *testing.T, noteArg markdownnotes.Note,
) *bool {
	called := false
	f.onPresentNote = func(n markdownnotes.Note) {
		if !reflect.DeepEqual(n, noteArg) {
			t.Errorf("Wrong note argument. Expected: %v. Actual: %v", noteArg, n)
		}
		called = true
	}
	return &called
}

func NewFakeNotePresenter() *FakeNotePresenter {
	return &FakeNotePresenter{
		onPresentNote: func(markdownnotes.Note) {},
	}
}
