package doubles

import (
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/cmd/server/presenters"
)

type HTTPNotePresenterSpy struct {
	*CreatedNotePresenterSpy
	*NotePresenterSpy
	*NoteListPresenterSpy
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
	CreateResponseWriterArg   http.ResponseWriter
	ReturnedHTTPNotePresenter *HTTPNotePresenterSpy
}

func NewHTTPNotePresenterFactorySpy() *HTTPNotePresenterFactorySpy {
	return &HTTPNotePresenterFactorySpy{}
}

func (s *HTTPNotePresenterFactorySpy) Create(w http.ResponseWriter) presenters.HTTPNotePresenter {
	s.CreateCalled = true
	s.CreateResponseWriterArg = w
	s.ReturnedHTTPNotePresenter = NewHTTPNotePresenterSpy()
	return s.ReturnedHTTPNotePresenter
}
