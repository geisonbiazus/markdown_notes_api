package notes

import (
	"errors"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
	"github.com/geisonbiazus/markdown_notes_api/internal/testing/mocks"
)

func TestCreateNoteUseCase(t *testing.T) {
	t.Run("It creates and stores a note", func(t *testing.T) {
		note := markdownnotes.Note{
			Title:   "Note Title",
			Content: "# Note Content",
		}

		storage := mocks.NewFakeNoteStorage()
		called := storage.ShouldReceiveSaveWithNoteAndReturn(t, note, nil)
		usecase := CreateNoteUseCase{storage}

		usecase.Run(note.Title, note.Content)

		if !*called {
			t.Errorf("It didn't store the note.")
		}
	})

	t.Run("It returns and error when there's and error on save", func(t *testing.T) {
		note := markdownnotes.Note{
			Title:   "Note Title",
			Content: "# Note Content",
		}

		expectedErr := errors.New("My error")
		storage := mocks.NewFakeNoteStorage()
		called := storage.ShouldReceiveSaveWithNoteAndReturn(t, note, expectedErr)
		usecase := CreateNoteUseCase{storage}

		err := usecase.Run(note.Title, note.Content)

		if err != expectedErr {
			t.Errorf(
				"It didn't return the correct error. Expected: %v. Actual: %v",
				expectedErr, err,
			)
		}

		if !*called {
			t.Errorf("It didn't store the note.")
		}
	})
}
