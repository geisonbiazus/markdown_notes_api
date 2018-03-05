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
		*doubles.NoteUseCaseSpy,
		*doubles.HTTPNotePresenterFactorySpy,
		*CreateNoteHandler,
		*httptest.ResponseRecorder,
	) {
		usecase := doubles.NewNoteUseCaseSpy()
		presenterFactory := doubles.NewHTTPNotePresenterFactorySpy()
		handler := NewCreateNoteHandler(usecase, presenterFactory)
		w := httptest.NewRecorder()

		return usecase, presenterFactory, handler, w
	}

	t.Run("Given a request title and content, it extracts the values and send them to use case", func(t *testing.T) {
		usecase, presenterFactory, handler, w := setup()

		title := "Title"
		content := "Content"

		r := createNoteRequest(`{"note": {"title": "` + title + `", "content": "` + content + `"}}`)

		handler.ServeHTTP(w, r)

		if presenterFactory.CreateResponseWritterArg != w {
			t.Errorf("It didn't create the presenter with response writer")
		}

		if usecase.CreateNoteTitleArg != title {
			t.Errorf("Expected: %v. Actual: %v", title, usecase.CreateNoteTitleArg)
		}

		if usecase.CreateNoteContentArg != content {
			t.Errorf("Expected: %v. Actual: %v", content, usecase.CreateNoteContentArg)
		}

		if usecase.CreateNotePresenterArg != presenterFactory.ReturnedHTTPNotePresenter {
			t.Errorf("Expected: %v. Actual: %v", presenterFactory.ReturnedHTTPNotePresenter, usecase.CreateNotePresenterArg)
		}
	})

	t.Run("Given a request with an invalid body JSON, ignores the body", func(t *testing.T) {
		usecase, presenterFactory, handler, w := setup()
		r := createNoteRequest("invalid JSON")

		handler.ServeHTTP(w, r)

		if usecase.CreateNoteTitleArg != "" {
			t.Errorf("Expected: %v. Actual: %v", "", usecase.CreateNoteTitleArg)
		}

		if usecase.CreateNoteContentArg != "" {
			t.Errorf("Expected: %v. Actual: %v", "", usecase.CreateNoteContentArg)
		}

		if usecase.CreateNotePresenterArg != presenterFactory.ReturnedHTTPNotePresenter {
			t.Errorf("Expected: %v. Actual: %v", presenterFactory.ReturnedHTTPNotePresenter, usecase.CreateNotePresenterArg)
		}
	})

	t.Run("When an error occurs in the use case, it presents internal server error", func(t *testing.T) {
		usecase, presenterFactory, handler, w := setup()
		usecase.CreateNoteErrorResult = errors.New("Error")

		r := createNoteRequest("")

		handler.ServeHTTP(w, r)

		if !presenterFactory.ReturnedHTTPNotePresenter.ServiceUnavailableCalled {
			t.Error("It didn't called presenter.ServiceUnavailable()")
		}
	})
}

func createNoteRequest(body string) *http.Request {
	return httptest.NewRequest("POST", "http://example.com/v1/notes", strings.NewReader(body))
}
