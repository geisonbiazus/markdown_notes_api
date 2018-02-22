package handlers

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

func TestListNotesHandler(t *testing.T) {
	t.Run("It runs the services passing a presenter", func(t *testing.T) {
		usecase := NewListNotesUseCaseSpy()
		presenterFactory := NewHTTPNoteListPresenterFactorySpy()
		handler := NewListNotesHandler(usecase, presenterFactory)
		r := httptest.NewRequest(http.MethodGet, "http://example.org/notes", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)

		if !reflect.DeepEqual(presenterFactory.CreateResponseWriterArg, w) {
			t.Errorf("Expected: %v. Actual: %v", w, presenterFactory.CreateResponseWriterArg)
		}

		if !reflect.DeepEqual(usecase.RunPresenterArg, presenterFactory.ReturnedHTTPNoteListPresenter) {
			t.Errorf("Expected: %v. Actual: %v", presenterFactory.ReturnedHTTPNoteListPresenter, usecase.RunPresenterArg)
		}
	})
}

type ListNotesUseCaseSpy struct {
	RunPresenterArg markdownnotes.NoteListPresenter
}

func NewListNotesUseCaseSpy() *ListNotesUseCaseSpy {
	return &ListNotesUseCaseSpy{}
}

func (s *ListNotesUseCaseSpy) Run(p markdownnotes.NoteListPresenter) {
	s.RunPresenterArg = p
}

type HTTPNoteListPresenterSpy struct{}

func (s *HTTPNoteListPresenterSpy) PresentNotes(n []markdownnotes.Note) {
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
