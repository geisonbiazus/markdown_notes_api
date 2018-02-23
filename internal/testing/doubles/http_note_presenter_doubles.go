package doubles

import (
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/cmd/server/presenters"
	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type HTTPNotePresenterSpy struct {
	*NotePresenterSpy
	ResponseWriter           http.ResponseWriter
	ServiceUnavailableCalled bool
}

func (s *HTTPNotePresenterSpy) ServiceUnavailable() {
	s.ServiceUnavailableCalled = true
}

func NewHTTPNotePresenterSpy() *HTTPNotePresenterSpy {
	return &HTTPNotePresenterSpy{}
}

type HTTPNotePresenterFactorySpy struct {
	CreateCalled              bool
	CreateResponseWritterArg  http.ResponseWriter
	ReturnedHTTPNotePresenter *HTTPNotePresenterSpy
}

func (s *HTTPNotePresenterFactorySpy) Create(w http.ResponseWriter) presenters.HTTPNotePresenter {
	s.CreateCalled = true
	s.CreateResponseWritterArg = w
	s.ReturnedHTTPNotePresenter = NewHTTPNotePresenterSpy()
	return s.ReturnedHTTPNotePresenter
}

func NewHTTPNotePresenterFactorySpy() *HTTPNotePresenterFactorySpy {
	return &HTTPNotePresenterFactorySpy{}
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

func (s *HTTPNoteListPresenterFactorySpy) Create(w http.ResponseWriter) presenters.HTTPNoteListPresenter {
	s.CreateResponseWriterArg = w
	return s.ReturnedHTTPNoteListPresenter
}
