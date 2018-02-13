package handlers

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/testing/doubles"
)

func TestCreateNoteHandler(t *testing.T) {
	t.Run("Given a request title and content, it extracts the values and send them to use case", func(t *testing.T) {
		usecase := doubles.NewCreateNoteUseCaseSpy()
		presenter := doubles.NewHTTPPresenterSpy()
		handler := NewCreateNoteHandler(usecase, presenter)

		title := "Title"
		content := "Content"

		body := strings.NewReader(`{"note": {"title": "` + title + `", "content": "` + content + `"}}`)
		r := httptest.NewRequest("POST", "http://example.com/v1/notes", body)
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, r)

		if presenter.ResponseWriter != w {
			t.Errorf("It didn't set the presenter response writer")
		}

		if usecase.RunTitleArg != title {
			t.Errorf("Expected: %v. Actual: %v", title, usecase.RunTitleArg)
		}

		if usecase.RunContentArg != content {
			t.Errorf("Expected: %v. Actual: %v", content, usecase.RunContentArg)
		}

		if usecase.RunPresenterArg != presenter {
			t.Errorf("Expected: %v. Actual: %v", presenter, usecase.RunPresenterArg)
		}
	})
}
