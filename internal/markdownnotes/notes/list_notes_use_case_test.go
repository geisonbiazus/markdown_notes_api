package notes

import (
	"reflect"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
	"github.com/geisonbiazus/markdown_notes_api/internal/testing/doubles"
)

func TestListNotesUseCase(t *testing.T) {
	setup := func() (
		*doubles.NoteStorageSpy,
		*ListNotesUseCase,
		*doubles.NoteListPresenterSpy,
	) {
		storage := doubles.NewNoteStorageSpy()
		usecase := NewListNotesUseCase(storage)
		presenter := doubles.NewNoteListPresenterSpy()
		return storage, usecase, presenter
	}

	t.Run("Given no notes, it presents an empty notes list", func(t *testing.T) {
		_, usecase, presenter := setup()
		usecase.Run(presenter)

		if !presenter.PresentNotesCalled {
			t.Error("It didn't call PresentNotes")
		}

		if len(presenter.PresentNotesNotesArg) > 1 {
			t.Error("It presented some notes when it shouldn't")
		}
	})

	t.Run("Given a saved note, it presentes a list containing it", func(t *testing.T) {
		storage, usecase, presenter := setup()

		notes := []markdownnotes.Note{
			markdownnotes.Note{Title: "Title", Content: "Content"},
		}

		storage.FindAllResult = notes
		usecase.Run(presenter)

		if !presenter.PresentNotesCalled {
			t.Error("It didn't call PresentNotes")
		}

		if !reflect.DeepEqual(presenter.PresentNotesNotesArg, notes) {
			t.Errorf("Expected: %v. Actual: %v", notes, presenter.PresentNotesNotesArg)
		}
	})

	t.Run("Given an error on find, it returns the error", func(t *testing.T) {
		_, usecase, presenter := setup()
		storage := doubles.NewErrorNoteSotorageStub()
		usecase.NoteStorage = storage

		err := usecase.Run(presenter)

		if err != storage.Error {
			t.Errorf("Expected: %v. Actual: %v", storage.Error, err)
		}
	})
}
