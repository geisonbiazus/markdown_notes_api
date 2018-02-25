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

			assertResponse(t, w, expectedBody, http.StatusCreated)
		})
	})
}
