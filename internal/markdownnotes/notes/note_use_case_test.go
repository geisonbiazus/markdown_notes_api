package notes

import (
	"errors"
	"reflect"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
	"github.com/geisonbiazus/markdown_notes_api/internal/testing/doubles"
)

func TestNoteUseCase(t *testing.T) {
	t.Run("CreateNote", func(t *testing.T) {
		setup := func() (
			markdownnotes.Note,
			*doubles.NoteStorageSpy,
			*doubles.NotePresenterSpy,
			*NoteUseCase,
		) {
			note := markdownnotes.Note{
				Title:   "Note Title",
				Content: "# Note Content",
			}
			storage := doubles.NewNoteStorageSpy()
			presenter := doubles.NewNotePresenterSpy()
			usecase := NewNoteUseCase(storage)

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

			usecase.CreateNote(note.Title, note.Content, presenter)

			if storage.SaveNoteArg != note {
				t.Errorf("Expected: %v. Actual: %v", note, storage.SaveNoteArg)
			}

			if presenter.PresentCreatedNoteNoteArg != savedNote {
				t.Errorf("Expected: %v. Actual: %v", savedNote, presenter.PresentCreatedNoteNoteArg)
			}
		})

		t.Run("Given an error occurs on create, it returns the error", func(t *testing.T) {
			note, _, presenter, usecase := setup()
			storage := doubles.NewErrorNoteSotorageStub()

			usecase.NoteStorage = storage

			err := usecase.CreateNote(note.Title, note.Content, presenter)

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

			usecase.CreateNote("", "", presenter)

			if !reflect.DeepEqual(presenter.PresentValidationErrorsErrsArg, errs) {
				t.Errorf("Expected: %v. Actual: %v", presenter.PresentValidationErrorsErrsArg, errs)
			}
		})
	})

	t.Run("ListNotes", func(t *testing.T) {
		setup := func() (
			*doubles.NoteStorageSpy,
			*NoteUseCase,
			*doubles.NotePresenterSpy,
		) {
			storage := doubles.NewNoteStorageSpy()
			usecase := NewNoteUseCase(storage)
			presenter := doubles.NewNotePresenterSpy()
			return storage, usecase, presenter
		}

		t.Run("Given no notes, it presents an empty notes list", func(t *testing.T) {
			_, usecase, presenter := setup()
			usecase.ListNotes(presenter)

			if !presenter.PresentNoteListCalled {
				t.Error("It didn't call PresentNoteList")
			}

			if len(presenter.PresentNoteListNotesArg) > 1 {
				t.Error("It presented some notes when it shouldn't")
			}
		})

		t.Run("Given a saved note, it presentes a list containing it", func(t *testing.T) {
			storage, usecase, presenter := setup()

			notes := []markdownnotes.Note{
				markdownnotes.Note{Title: "Title", Content: "Content"},
			}

			storage.FindAllResult = notes
			usecase.ListNotes(presenter)

			if !presenter.PresentNoteListCalled {
				t.Error("It didn't call PresentNoteList")
			}

			if !reflect.DeepEqual(presenter.PresentNoteListNotesArg, notes) {
				t.Errorf("Expected: %v. Actual: %v", notes, presenter.PresentNoteListNotesArg)
			}
		})

		t.Run("Given an error on find, it returns the error", func(t *testing.T) {
			_, usecase, presenter := setup()
			storage := doubles.NewErrorNoteSotorageStub()
			usecase.NoteStorage = storage

			err := usecase.ListNotes(presenter)

			if err != storage.Error {
				t.Errorf("Expected: %v. Actual: %v", storage.Error, err)
			}
		})
	})

	t.Run("ShowNote", func(t *testing.T) {
		setup := func() (
			*doubles.NoteStorageSpy,
			*NoteUseCase,
			*doubles.NotePresenterSpy,
		) {
			storage := doubles.NewNoteStorageSpy()
			usecase := NewNoteUseCase(storage)
			presenter := doubles.NewNotePresenterSpy()
			return storage, usecase, presenter
		}

		t.Run("Given an existent note ID, it fetchs and presents the note", func(t *testing.T) {
			storage, usecase, presenter := setup()

			note := markdownnotes.Note{ID: 1, Title: "Title"}
			storage.FindByIDNoteResult = note

			usecase.ShowNote(note.ID, presenter)

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

			usecase.ShowNote(noteID, presenter)

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

			err := usecase.ShowNote(noteID, presenter)

			if err != storage.Error {
				t.Errorf("Expected: %v. Actual: %v", err, storage.Error)
			}
		})
	})

	t.Run("UpdateNote", func(t *testing.T) {
		setup := func() (
			*NoteUseCase,
			*doubles.NoteStorageSpy,
			*doubles.NotePresenterSpy,
			markdownnotes.Note,
		) {
			storage := doubles.NewNoteStorageSpy()
			usecase := NewNoteUseCase(storage)
			presenter := doubles.NewNotePresenterSpy()

			savedNote := markdownnotes.Note{
				ID:      42,
				Title:   "Title",
				Content: "Content",
			}

			storage.FindByIDNoteResult = savedNote

			return usecase, storage, presenter, savedNote
		}

		t.Run("Given an existing note, it updates and presents it", func(t *testing.T) {
			usecase, storage, presenter, savedNote := setup()

			expectedNote := markdownnotes.Note{
				ID:      savedNote.ID,
				Title:   "New Title",
				Content: "New Content",
			}

			storage.SaveNoteResult = expectedNote

			usecase.UpdateNote(savedNote.ID, expectedNote.Title, expectedNote.Content, presenter)

			if storage.SaveNoteArg != expectedNote {
				t.Errorf("Expected: %v. Actual: %v", expectedNote, storage.SaveNoteArg)
			}

			if presenter.PresentUpdatedNoteNoteArg != expectedNote {
				t.Errorf("Expected: %v. Actual: %v", expectedNote, presenter.PresentUpdatedNoteNoteArg)
			}
		})

		t.Run("Given the note isn'saved, it presents NotFound", func(t *testing.T) {
			usecase, storage, presenter, savedNote := setup()
			storage.FindByIDNoteResult = markdownnotes.NoNote

			usecase.UpdateNote(savedNote.ID, "title", "content", presenter)

			if !presenter.NotFoundCalled {
				t.Error("It didn't call NotFound")
			}
		})

		t.Run("Given invalid params, it doesn't save the note and presents the errors", func(t *testing.T) {
			usecase, storage, presenter, savedNote := setup()
			usecase.UpdateNote(savedNote.ID, "", "", presenter)

			errs := []markdownnotes.ValidationError{
				markdownnotes.ValidationError{
					Field:   "title",
					Message: "Can't be blank",
					Type:    "REQUIRED",
				},
			}

			if !reflect.DeepEqual(presenter.PresentValidationErrorsErrsArg, errs) {
				t.Errorf("Expected: %v. Actual: %v", presenter.PresentValidationErrorsErrsArg, errs)
			}

			if storage.SaveCalled {
				t.Error("It called save when it shouldn't")
			}
		})

		t.Run("If an error occurs on find, it returns the error", func(t *testing.T) {
			usecase, _, presenter, savedNote := setup()
			storage := doubles.NewErrorNoteSotorageStub()
			usecase.NoteStorage = storage

			err := usecase.UpdateNote(savedNote.ID, "title", "content", presenter)

			if err != storage.Error {
				t.Errorf("Expected: %v. Actual: %v", storage.Error, err)
			}
		})

		t.Run("If an error occurs on save, it returns the error", func(t *testing.T) {
			usecase, storage, presenter, savedNote := setup()
			storage.SaveNoteErrorResult = errors.New("Some error")

			err := usecase.UpdateNote(savedNote.ID, "title", "content", presenter)

			if err != storage.SaveNoteErrorResult {
				t.Errorf("Expected: %v. Actual: %v", storage.SaveNoteErrorResult, err)
			}
		})
	})
}
