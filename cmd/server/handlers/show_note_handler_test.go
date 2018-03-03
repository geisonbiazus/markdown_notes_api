package handlers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/testing/doubles"
	"github.com/julienschmidt/httprouter"
)

func TestShowNoteHandler(t *testing.T) {
	setup := func() (
		*httptest.ResponseRecorder,
		*doubles.HTTPNotePresenterFactorySpy,
		*doubles.ShowNoteUseCaseSpy,
		*ShowNoteHandler,
	) {
		w := httptest.NewRecorder()
		presenterFactory := doubles.NewHTTPNotePresenterFactorySpy()
		usecase := doubles.NewShowNoteUseCaseSpy()
		handler := NewShowNoteHandler(usecase, presenterFactory)

		return w, presenterFactory, usecase, handler
	}

	newRequest := func(noteID string) *http.Request {
		r := httptest.NewRequest(http.MethodGet, "http://example.org/api/v1/notes/"+noteID, nil)
		params := httprouter.Params{httprouter.Param{"id", noteID}}
		ctx := context.WithValue(r.Context(), httprouter.ParamsKey, params)
		r = r.WithContext(ctx)

		return r
	}

	t.Run("Given an id param, it pass the ID and a presenter to the usecase", func(t *testing.T) {
		w, presenterFactory, usecase, handler := setup()
		noteID := 42
		r := newRequest(strconv.Itoa(noteID))

		handler.ServeHTTP(w, r)

		if usecase.RunNoteIDArg != noteID {
			t.Errorf("Expected: %v. Actual: %v", noteID, usecase.RunNoteIDArg)
		}

		if usecase.RunPresenterArg != presenterFactory.ReturnedHTTPNotePresenter {
			t.Errorf("Expected: %v. Actual: %v", presenterFactory.ReturnedHTTPNotePresenter, usecase.RunPresenterArg)
		}
	})

	t.Run("Given an invalid id param, it passes 0 to usecase", func(t *testing.T) {
		w, _, usecase, handler := setup()
		r := newRequest("invalid")

		handler.ServeHTTP(w, r)

		if usecase.RunNoteIDArg != 0 {
			t.Errorf("Expected: %v. Actual: %v", 0, usecase.RunNoteIDArg)
		}
	})

	t.Run("When an error occurs on UseCase, it presents ServiceUnavailable", func(t *testing.T) {
		w, presenterFactory, usecase, handler := setup()
		r := newRequest("42")
		usecase.RunErrorResult = errors.New("error")

		handler.ServeHTTP(w, r)

		if !presenterFactory.ReturnedHTTPNotePresenter.ServiceUnavailableCalled {
			t.Error("It didn't call ServiceUnavailable")
		}
	})
}
