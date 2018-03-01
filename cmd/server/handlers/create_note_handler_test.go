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
		*doubles.HTTPNotePresenterFactorySpy,
		*CreateNoteHandler,
		*httptest.ResponseRecorder,
	) {
		usecase := doubles.NewCreateNoteUseCaseSpy()
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

		if usecase.RunTitleArg != title {
			t.Errorf("Expected: %v. Actual: %v", title, usecase.RunTitleArg)
		}

		if usecase.RunContentArg != content {
			t.Errorf("Expected: %v. Actual: %v", content, usecase.RunContentArg)
		}

		if usecase.RunPresenterArg != presenterFactory.ReturnedHTTPNotePresenter {
			t.Errorf("Expected: %v. Actual: %v", presenterFactory.ReturnedHTTPNotePresenter, usecase.RunPresenterArg)
		}
	})

	t.Run("Given a request with an invalid body JSON, ignores the body", func(t *testing.T) {
		usecase, presenterFactory, handler, w := setup()
		r := createNoteRequest("invalid JSON")

		handler.ServeHTTP(w, r)

		if usecase.RunTitleArg != "" {
			t.Errorf("Expected: %v. Actual: %v", "", usecase.RunTitleArg)
		}

		if usecase.RunContentArg != "" {
			t.Errorf("Expected: %v. Actual: %v", "", usecase.RunContentArg)
		}

		if usecase.RunPresenterArg != presenterFactory.ReturnedHTTPNotePresenter {
			t.Errorf("Expected: %v. Actual: %v", presenterFactory.ReturnedHTTPNotePresenter, usecase.RunPresenterArg)
		}
	})

	t.Run("When an error occurs in the use case, it presents internal server error", func(t *testing.T) {
		usecase, presenterFactory, handler, w := setup()
		usecase.RunErrorResult = errors.New("Error")

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
