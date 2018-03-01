package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/testing/doubles"
	"github.com/julienschmidt/httprouter"
)

func TestShowNoteHandler(t *testing.T) {
	setup := func() (int,
		*httptest.ResponseRecorder,
		*http.Request,
		httprouter.Params,
		*doubles.HTTPNotePresenterFactorySpy,
		*doubles.ShowNoteUseCaseSpy,
		*ShowNoteHandler,
	) {
		noteID := 42
		noteIDString := strconv.Itoa(noteID)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "http://example.org/api/v1/notes/"+noteIDString, nil)
		params := httprouter.Params{httprouter.Param{"id", noteIDString}}
		presenterFactory := doubles.NewHTTPNotePresenterFactorySpy()
		usecase := doubles.NewShowNoteUseCaseSpy()
		handler := NewShowNoteHandler(usecase, presenterFactory)

		return noteID, w, r, params, presenterFactory, usecase, handler
	}

	t.Run("Given an id param, it pass the ID and a presenter to the usecase", func(t *testing.T) {
		noteID, w, r, params, presenterFactory, usecase, handler := setup()

		handler.ServeHTTP(w, r, params)

		if usecase.RunNoteIDArg != noteID {
			t.Errorf("Expected: %v. Actual: %v", noteID, usecase.RunNoteIDArg)
		}

		if usecase.RunPresenterArg != presenterFactory.ReturnedHTTPNotePresenter {
			t.Errorf("Expected: %v. Actual: %v", presenterFactory.ReturnedHTTPNotePresenter, usecase.RunPresenterArg)
		}
	})

	t.Run("Given an invalid id param, it passes 0 to usecase", func(t *testing.T) {
		_, w, r, params, _, usecase, handler := setup()
		params[0].Value = "invalid"

		handler.ServeHTTP(w, r, params)

		if usecase.RunNoteIDArg != 0 {
			t.Errorf("Expected: %v. Actual: %v", 0, usecase.RunNoteIDArg)
		}
	})

	t.Run("When an error occurs on UseCase, it presents ServiceUnavailable", func(t *testing.T) {
		_, w, r, params, presenterFactory, usecase, handler := setup()
		usecase.RunErrorResult = errors.New("error")

		handler.ServeHTTP(w, r, params)

		if !presenterFactory.ReturnedHTTPNotePresenter.ServiceUnavailableCalled {
			t.Error("It didn't call ServiceUnavailable")
		}
	})
}
