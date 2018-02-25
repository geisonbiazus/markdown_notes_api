package notes

import (
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
	"github.com/geisonbiazus/markdown_notes_api/internal/testing/doubles"
)

func TestShowNoteUseCase(t *testing.T) {
	setup := func() (
		*doubles.NoteStorageSpy,
		*ShowNoteUseCase,
		*doubles.NotePresenterSpy,
	) {
		storage := doubles.NewNoteStorageSpy()
		usecase := NewShowNoteUseCase(storage)
		presenter := doubles.NewNotePresenterSpy()
		return storage, usecase, presenter
	}

	t.Run("Given an existent note ID, it fetchs and presents the note", func(t *testing.T) {
		storage, usecase, presenter := setup()

		note := markdownnotes.Note{ID: 1, Title: "Title"}
		storage.FindByIDNoteResult = note

		usecase.Run(note.ID, presenter)

		if storage.FindByIDIDArg != note.ID {
			t.Errorf("Expected: %v. Actual: %v", note.ID, storage.FindByIDIDArg)
		}

		if presenter.PresentNoteNoteArg != note {
			t.Errorf("Expected: %v. Actual: %v", note, presenter.PresentNoteNoteArg)
		}
	})

	t.Run("Given a non existing note ID, it presents note not found", func(t *testing.T) {
		storage, usecase, presenter := setup()

		noteID := 42

		usecase.Run(noteID, presenter)

		if storage.FindByIDIDArg != noteID {
			t.Errorf("Expected: %v. Actual: %v", noteID, storage.FindByIDIDArg)
		}

		if !presenter.NotFoundCalled {
			t.Error("It didn't call NotFound")
		}
	})

	t.Run("When an error occur on find, it returns the error", func(t *testing.T) {
		_, usecase, presenter := setup()
		storage := doubles.NewErrorNoteSotorageStub()
		usecase.NoteStorage = storage

		noteID := 42

		err := usecase.Run(noteID, presenter)

		if err != storage.Error {
			t.Errorf("Expected: %v. Actual: %v", err, storage.Error)
		}
	})
}
