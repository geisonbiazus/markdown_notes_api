package notes

import (
	"reflect"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
	"github.com/geisonbiazus/markdown_notes_api/internal/testing/doubles"
)

func TestCreateNoteUseCase(t *testing.T) {
	setup := func() (
		markdownnotes.Note,
		*doubles.NoteStorageSpy,
		*doubles.NotePresenterSpy,
		*CreateNoteUseCase,
	) {
		note := markdownnotes.Note{
			Title:   "Note Title",
			Content: "# Note Content",
		}
		storage := doubles.NewNoteStorageSpy()
		presenter := doubles.NewNotePresenterSpy()
		usecaseFactory := &CreateNoteUseCaseFactory{storage}
		usecase := usecaseFactory.Create(presenter)

		return note, storage, presenter, usecase
	}

	t.Run("Given valid arguments, it creates and presents a note", func(t *testing.T) {
		note, storage, presenter, usecase := setup()
		savedNote := markdownnotes.Note{
			ID:      1,
			Title:   note.Title,
			Content: note.Content,
		}

		storage.SaveNoteResult = savedNote

		usecase.Run(note.Title, note.Content)

		if storage.SaveNoteArg() != note {
			t.Errorf("Expected: %v. Actual: %v", note, storage.SaveNoteArg())
		}

		if presenter.PresentNoteNoteArg != savedNote {
			t.Errorf("Expected: %v. Actual: %v", savedNote, presenter.PresentNoteNoteArg)
		}
	})

	t.Run("Given an error occurs on create, it returns the error", func(t *testing.T) {
		note, _, _, usecase := setup()
		storage := doubles.NewErrorNoteSotorageStub()

		usecase.NoteStorage = storage

		err := usecase.Run(note.Title, note.Content)

		if storage.Error != err {
			t.Errorf("Expected: %v. Actual: %v", storage.Error, err)
		}
	})

	t.Run("Given an empty title, it validates note and presents the error", func(t *testing.T) {
		_, _, presenter, usecase := setup()

		errs := []markdownnotes.ValidationError{
			markdownnotes.ValidationError{
				Field:   "title",
				Message: "Can't be blank",
				Type:    "REQUIRED",
			},
		}

		usecase.Run("", "")

		if !reflect.DeepEqual(presenter.PresentErrorErrsArg, errs) {
			t.Errorf("Expected: %v. Actual: %v", presenter.PresentErrorErrsArg, errs)
		}
	})
}
