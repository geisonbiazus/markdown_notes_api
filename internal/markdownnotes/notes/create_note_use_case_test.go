package notes

import (
	"errors"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
	"github.com/geisonbiazus/markdown_notes_api/internal/testing/mocks"
)

func TestCreateNoteUseCase(t *testing.T) {
	setup := func() (
		markdownnotes.Note,
		*mocks.FakeNoteStorage,
		*mocks.FakeNotePresenter,
		CreateNoteUseCase,
	) {
		note := markdownnotes.Note{
			Title:   "Note Title",
			Content: "# Note Content",
		}
		storage := mocks.NewFakeNoteStorage()
		presenter := mocks.NewFakeNotePresenter()
		usecase := CreateNoteUseCase{storage, presenter}

		return note, storage, presenter, usecase
	}

	t.Run("It creates a note, stores and presents it", func(t *testing.T) {
		note, storage, presenter, usecase := setup()
		savedNote := markdownnotes.Note{
			ID:      1,
			Title:   note.Title,
			Content: note.Content,
		}

		saveCalled := storage.ShouldReceiveSaveWithNoteAndReturn(t, note, savedNote, nil)
		presentCalled := presenter.ShouldReceivePresentNoteWithNote(t, savedNote)
		usecase.Run(note.Title, note.Content)

		if !*saveCalled {
			t.Errorf("It didn't store the note.")
		}

		if !*presentCalled {
			t.Errorf("It didn't present the note.")
		}
	})

	t.Run("It returns an error when there's and error on save", func(t *testing.T) {
		note, storage, _, usecase := setup()

		expectedErr := errors.New("My error")

		storage.ShouldReceiveSaveWithNoteAndReturn(t, note, note, expectedErr)
		err := usecase.Run(note.Title, note.Content)

		if err != expectedErr {
			t.Errorf("Expected: %v. Actual: %v", expectedErr, err)
		}
	})

	t.Run("It validates note with an empty title and Presents the error", func(t *testing.T) {
		_, storage, presenter, usecase := setup()

		errs := []markdownnotes.ValidationError{
			markdownnotes.ValidationError{
				Field:   "title",
				Message: "Can't be blank",
				Type:    "REQUIRED",
			},
		}

		storage.ShouldNotReceiveSave(t)
		called := presenter.ShouldReceivePresentErrorsWithErrors(t, errs)

		usecase.Run("", "")

		if !*called {
			t.Errorf("It didn't present the errors.")
		}

	})
}
