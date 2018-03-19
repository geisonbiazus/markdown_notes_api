package handlers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/testing/doubles"
	"github.com/julienschmidt/httprouter"
)

type updateNoteHandlerFixture struct {
	usecase          *doubles.NoteUseCaseSpy
	presenterFactory *doubles.NotePresenterFactorySpy
	handler          *UpdateNoteHandler
	w                *httptest.ResponseRecorder
	r                *http.Request
	id               int
	title            string
	content          string
}

func TestUpdateNoteHandler(t *testing.T) {
	setup := func() *updateNoteHandlerFixture {
		usecase := doubles.NewNoteUseCaseSpy()
		presenterFactory := doubles.NewNotePresenterFactorySpy()

		handler := NewUpdateNoteHandler(usecase, presenterFactory)

		w := httptest.NewRecorder()

		id, title, content := 1, "New Title", "New Content"

		body := `{"note":{"title":"` + title + `","content":"` + content + `"}}`
		r := httptest.NewRequest(http.MethodPut, "http://example.org/api/v1/notes/1", strings.NewReader(body))

		params := httprouter.Params{httprouter.Param{Key: "id", Value: strconv.Itoa(id)}}
		r = r.WithContext(context.WithValue(r.Context(), httprouter.ParamsKey, params))

		return &updateNoteHandlerFixture{
			usecase:          usecase,
			presenterFactory: presenterFactory,
			handler:          handler,
			w:                w,
			r:                r,
			id:               id,
			title:            title,
			content:          content,
		}
	}

	t.Run("Given the correct params, it updates the note", func(t *testing.T) {
		f := setup()

		f.handler.ServeHTTP(f.w, f.r)

		if f.usecase.UpdateNoteIDArg != f.id {
			t.Errorf("\nExpected: %v\n  Actual: %v", f.id, f.usecase.UpdateNoteIDArg)
		}

		if f.usecase.UpdateNoteTitleArg != f.title {
			t.Errorf("\nExpected: %v\n  Actual: %v", f.title, f.usecase.UpdateNoteTitleArg)
		}

		if f.usecase.UpdateNoteContentArg != f.content {
			t.Errorf("\nExpected: %v\n  Actual: %v", f.content, f.usecase.UpdateNoteContentArg)
		}

		if f.usecase.UpdateNotePresenterArg != f.presenterFactory.ReturnedNotePresenter {
			t.Errorf("\nExpected: %v\n  Actual: %v", f.presenterFactory.ReturnedNotePresenter, f.usecase.UpdateNotePresenterArg)
		}
	})

	t.Run("When an error occurs, it presents the error", func(t *testing.T) {
		f := setup()
		err := errors.New("Some error")
		f.usecase.UpdateNoteErrorResult = err

		f.handler.ServeHTTP(f.w, f.r)

		if !f.presenterFactory.ReturnedNotePresenter.PresentErrorCalled {
			t.Error("It should present the error")
		}

		if f.presenterFactory.ReturnedNotePresenter.PresentErrorErrorArg != err {
			t.Errorf("\nExpected: %v\n  Actual: %v", err, f.presenterFactory.ReturnedNotePresenter.PresentErrorErrorArg)
		}
	})
}
