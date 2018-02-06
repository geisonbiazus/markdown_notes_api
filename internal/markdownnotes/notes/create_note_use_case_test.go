package notes

import (
	"reflect"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

func TestCreateNoteUseCase(t *testing.T) {
	t.Run("It creates and stores a note", func(t *testing.T) {
		title := "Note Title"
		content := "# Note Content"

		note := markdownnotes.Note{
			Title:   title,
			Content: content,
		}

		storage := newFakeNoteStorage()
		called := false

		storage.onSave = func(n markdownnotes.Note) {
			if !reflect.DeepEqual(n, note) {
				t.Errorf("It didn't store the properly. Expected: %v. Actual: %v", note, n)
			}
			called = true
		}

		usecase := CreateNoteUseCase{
			NoteStorage: storage,
		}

		usecase.Run(title, content)

		if !called {
			t.Errorf("It didn't store the note.")
		}
	})
}

type fakeNoteStorage struct {
	onSave func(markdownnotes.Note)
}

func (f *fakeNoteStorage) Save(n markdownnotes.Note) {
	f.onSave(n)
}

func newFakeNoteStorage() *fakeNoteStorage {
	return new(fakeNoteStorage)
}
