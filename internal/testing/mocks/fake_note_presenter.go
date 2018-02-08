package mocks

import (
	"reflect"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type FakeNotePresenter struct {
	onPresentNote  func(note markdownnotes.Note)
	onPresentError func(errs []markdownnotes.ValidationError)
}

func (f *FakeNotePresenter) PresentNote(note markdownnotes.Note) {
	if f.onPresentNote != nil {
		f.onPresentNote(note)
	}
}

func (f *FakeNotePresenter) PresentError(errs []markdownnotes.ValidationError) {
	if f.onPresentError != nil {
		f.onPresentError(errs)
	}
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

func (f *FakeNotePresenter) ShouldReceivePresentErrorsWithErrors(
	t *testing.T, errsArg []markdownnotes.ValidationError,
) *bool {
	called := false
	f.onPresentError = func(e []markdownnotes.ValidationError) {
		if !reflect.DeepEqual(e, errsArg) {
			t.Errorf("Wrong err argument. Expected: %v. Actual: %v", errsArg, e)
		}
		called = true
	}
	return &called
}

func NewFakeNotePresenter() *FakeNotePresenter {
	return new(FakeNotePresenter)
}
