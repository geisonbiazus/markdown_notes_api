package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

func TestListNotesHandler(t *testing.T) {
	setup := func() (
		*ListNotesUseCaseSpy,
		*HTTPNoteListPresenterFactorySpy,
		*ListNotesHandler,
		*http.Request,
		*httptest.ResponseRecorder,
	) {
		usecase := NewListNotesUseCaseSpy()
		presenterFactory := NewHTTPNoteListPresenterFactorySpy()
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

type ListNotesUseCaseSpy struct {
	RunPresenterArg markdownnotes.NoteListPresenter
	RunErrorResult  error
}

func NewListNotesUseCaseSpy() *ListNotesUseCaseSpy {
	return &ListNotesUseCaseSpy{}
}

func (s *ListNotesUseCaseSpy) Run(p markdownnotes.NoteListPresenter) error {
	s.RunPresenterArg = p
	return s.RunErrorResult
}

type HTTPNoteListPresenterSpy struct {
	ServiceUnavailableCalled bool
}

func (s *HTTPNoteListPresenterSpy) PresentNotes(n []markdownnotes.Note) {
}

func (s *HTTPNoteListPresenterSpy) ServiceUnavailable() {
	s.ServiceUnavailableCalled = true
}

type HTTPNoteListPresenterFactorySpy struct {
	ReturnedHTTPNoteListPresenter *HTTPNoteListPresenterSpy
	CreateResponseWriterArg       http.ResponseWriter
}

func NewHTTPNoteListPresenterFactorySpy() *HTTPNoteListPresenterFactorySpy {
	presenterSpy := &HTTPNoteListPresenterSpy{}
	return &HTTPNoteListPresenterFactorySpy{
		ReturnedHTTPNoteListPresenter: presenterSpy,
	}
}

func (s *HTTPNoteListPresenterFactorySpy) Create(w http.ResponseWriter) HTTPNoteListPresenter {
	s.CreateResponseWriterArg = w
	return s.ReturnedHTTPNoteListPresenter
}
