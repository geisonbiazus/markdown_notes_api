package notes

import (
	"errors"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
	"github.com/geisonbiazus/markdown_notes_api/internal/testing/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateNoteUseCase(t *testing.T) {
	setup := func() (
		markdownnotes.Note,
		*mocks.NoteStorageMock,
		*mocks.NotePresenterMock,
		*CreateNoteUseCase,
	) {
		note := markdownnotes.Note{
			Title:   "Note Title",
			Content: "# Note Content",
		}
		storage := mocks.NewNoteStorageMock()
		presenter := mocks.NewNotePresenterMock()
		usecaseFactory := &CreateNoteUseCaseFactory{storage}
		usecase := usecaseFactory.Create(presenter)

		return note, storage, presenter, usecase
	}

	t.Run("It creates a note, stores and presents it", func(t *testing.T) {
		note, storage, presenter, usecase := setup()
		savedNote := markdownnotes.Note{
			ID:      1,
			Title:   note.Title,
			Content: note.Content,
		}

		storage.On("Save", note).Return(savedNote, nil)
		presenter.On("PresentNote", savedNote)

		usecase.Run(note.Title, note.Content)

		storage.AssertExpectations(t)
		presenter.AssertExpectations(t)
	})

	t.Run("It returns an error when there's and error on save", func(t *testing.T) {
		note, storage, _, usecase := setup()

		expectedErr := errors.New("My error")

		storage.On("Save", note).Return(note, expectedErr)

		err := usecase.Run(note.Title, note.Content)

		assert.Equal(t, expectedErr, err)
	})

	t.Run("It validates note with an empty title and Presents the error", func(t *testing.T) {
		_, _, presenter, usecase := setup()

		errs := []markdownnotes.ValidationError{
			markdownnotes.ValidationError{
				Field:   "title",
				Message: "Can't be blank",
				Type:    "REQUIRED",
			},
		}

		presenter.On("PresentError", errs)

		usecase.Run("", "")

		presenter.AssertExpectations(t)
	})
}
