package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/testing/doubles"
)

func TestCreateNoteHandler(t *testing.T) {
	setup := func() (
		*doubles.CreateNoteUseCaseSpy,
		*doubles.HTTPNotePresenterSpy,
		*CreateNoteHandler,
		*httptest.ResponseRecorder,
	) {
		usecase := doubles.NewCreateNoteUseCaseSpy()
		presenter := doubles.NewHTTPPresenterSpy()
		handler := NewCreateNoteHandler(usecase, presenter)
		w := httptest.NewRecorder()

		return usecase, presenter, handler, w
	}

	t.Run("Given a request title and content, it extracts the values and send them to use case", func(t *testing.T) {
		usecase, presenter, handler, w := setup()

		title := "Title"
		content := "Content"

		r := createNoteRequest(`{"note": {"title": "` + title + `", "content": "` + content + `"}}`)

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

	t.Run("Given a request with an invalid body JSON, ignores the body", func(t *testing.T) {
		usecase, presenter, handler, w := setup()
		r := createNoteRequest("invalid JSON")

		handler.ServeHTTP(w, r)

		if usecase.RunTitleArg != "" {
			t.Errorf("Expected: %v. Actual: %v", "", usecase.RunTitleArg)
		}

		if usecase.RunContentArg != "" {
			t.Errorf("Expected: %v. Actual: %v", "", usecase.RunContentArg)
		}

		if usecase.RunPresenterArg != presenter {
			t.Errorf("Expected: %v. Actual: %v", presenter, usecase.RunPresenterArg)
		}
	})

	t.Run("When an error occur in use case, it presents internal server error", func(t *testing.T) {
		usecase, presenter, handler, w := setup()
		usecase.RunErrorResult = errors.New("Error")

		r := createNoteRequest("")

		handler.ServeHTTP(w, r)

		if !presenter.ServiceUnavailableCalled {
			t.Error("It didn't called presenter.ServiceUnavailable()")
		}
	})
}

func createNoteRequest(body string) *http.Request {
	return httptest.NewRequest("POST", "http://example.com/v1/notes", strings.NewReader(body))
}
