package notes

import (
	"errors"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
	"github.com/geisonbiazus/markdown_notes_api/internal/testing/mocks"
)

func TestCreateNoteUseCase(t *testing.T) {
	setup := func() (markdownnotes.Note, *mocks.FakeNoteStorage, CreateNoteUseCase) {
		note := markdownnotes.Note{
			Title:   "Note Title",
			Content: "# Note Content",
		}
		storage := mocks.NewFakeNoteStorage()
		usecase := CreateNoteUseCase{storage}

		return note, storage, usecase
	}

	t.Run("It creates and stores a note", func(t *testing.T) {
		note, storage, usecase := setup()

		called := storage.ShouldReceiveSaveWithNoteAndReturn(t, note, nil)
		usecase.Run(note.Title, note.Content)

		if !*called {
			t.Errorf("It didn't store the note.")
		}
	})

	t.Run("It returns an error when there's and error on save", func(t *testing.T) {
		note, storage, usecase := setup()

		expectedErr := errors.New("My error")

		storage.ShouldReceiveSaveWithNoteAndReturn(t, note, expectedErr)
		err := usecase.Run(note.Title, note.Content)

		if err != expectedErr {
			t.Errorf("Expected: %v. Actual: %v", expectedErr, err)
		}
	})
}
