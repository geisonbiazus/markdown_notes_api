package presenters

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

func TestNoteJSONPresenter(t *testing.T) {
	setup := func() (
		*httptest.ResponseRecorder,
		*NoteJSONPresenter,
	) {
		w := httptest.NewRecorder()
		factory := NewNoteJSONPresenterFactory()
		presenter := factory.Create(w).(*NoteJSONPresenter)

		return w, presenter
	}

	t.Run("PresentNote", func(t *testing.T) {
		t.Run("Given a note, it renders the note data as a JSON", func(t *testing.T) {
			w, presenter := setup()

			note := markdownnotes.Note{
				ID:      1,
				Title:   "Title",
				Content: "Content",
			}

			presenter.PresentNote(note)

			expectedBody := []byte(fmt.Sprintf(`{"id":%d,"title":"%s","content":"%s"}`, note.ID, note.Title, note.Content))

			assertResponse(t, w, expectedBody, http.StatusOK)
		})
	})

	t.Run("PresentCreatedNote", func(t *testing.T) {
		t.Run("Given a note, it renders the note data as a JSON", func(t *testing.T) {
			w, presenter := setup()

			note := markdownnotes.Note{
				ID:      1,
				Title:   "Title",
				Content: "Content",
			}

			presenter.PresentCreatedNote(note)

			expectedBody := []byte(fmt.Sprintf(`{"id":%d,"title":"%s","content":"%s"}`, note.ID, note.Title, note.Content))

			assertResponse(t, w, expectedBody, http.StatusCreated)
		})
	})

	t.Run("PresentNotes", func(t *testing.T) {
		t.Run("Given a list of notes, it renders the notes data as a JSON", func(t *testing.T) {
			w, presenter := setup()

			note1 := markdownnotes.Note{ID: 1, Title: "Title", Content: "Content"}
			note2 := markdownnotes.Note{ID: 2, Title: "Title 2", Content: "Content 2"}

			presenter.PresentNotes([]markdownnotes.Note{note1, note2})

			expectedBody := []byte(
				fmt.Sprintf(
					`[{"id":%d,"title":"%s","content":"%s"},{"id":%d,"title":"%s","content":"%s"}]`,
					note1.ID, note1.Title, note1.Content, note2.ID, note2.Title, note2.Content,
				),
			)

			assertResponse(t, w, expectedBody, http.StatusOK)
		})
	})
}
