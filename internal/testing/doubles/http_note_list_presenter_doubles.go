package doubles

import (
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/cmd/server/presenters"
)

type HTTPNoteListPresenterSpy struct {
	*NoteListPresenterSpy
	ServiceUnavailableCalled bool
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
