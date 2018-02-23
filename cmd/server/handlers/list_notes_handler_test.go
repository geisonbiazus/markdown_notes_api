package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/testing/doubles"
)

func TestListNotesHandler(t *testing.T) {
	setup := func() (
		*doubles.ListNotesUseCaseSpy,
		*doubles.HTTPNoteListPresenterFactorySpy,
		*ListNotesHandler,
		*http.Request,
		*httptest.ResponseRecorder,
	) {
		usecase := doubles.NewListNotesUseCaseSpy()
		presenterFactory := doubles.NewHTTPNoteListPresenterFactorySpy()
		handler := NewListNotesHandler(usecase, presenterFactory)
		r := httptest.NewRequest(http.MethodGet, "http://example.org/notes", nil)
		w := httptest.NewRecorder()
		return usecase, presenterFactory, handler, r, w
	}

	t.Run("It runs the services passing a presenter", func(t *testing.T) {
		usecase, presenterFactory, handler, r, w := setup()
		handler.ServeHTTP(w, r)

		if !reflect.DeepEqual(presenterFactory.CreateResponseWriterArg, w) {
			t.Errorf("Expected: %v. Actual: %v", w, presenterFactory.CreateResponseWriterArg)
		}

		if !reflect.DeepEqual(usecase.RunPresenterArg, presenterFactory.ReturnedHTTPNoteListPresenter) {
			t.Errorf("Expected: %v. Actual: %v", presenterFactory.ReturnedHTTPNoteListPresenter, usecase.RunPresenterArg)
		}
	})

	t.Run("When an error is returned from UseCase, it presents service unavailable", func(t *testing.T) {
		usecase, presenterFactory, handler, r, w := setup()
		usecase.RunErrorResult = errors.New("Error")
		handler.ServeHTTP(w, r)

		presenter := presenterFactory.ReturnedHTTPNoteListPresenter

		if !presenter.ServiceUnavailableCalled {
			t.Error("It didn't call ServiceUnavailable")
		}
	})
}
