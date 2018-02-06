package notes

import (
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
		called := storage.ShouldReceiveSaveWithNote(t, note)
		usecase := CreateNoteUseCase{storage}

		usecase.Run(note.Title, note.Content)

		if !*called {
			t.Errorf("It didn't store the note.")
		}
	})
}
