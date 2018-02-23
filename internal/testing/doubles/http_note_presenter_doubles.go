package doubles

import (
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/cmd/server/presenters"
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
