package doubles

import (
	"net/http"
)

type HTTPNotePresenterSpy struct {
	*NotePresenterSpy
	ResponseWriter            http.ResponseWriter
	InternalServerErrorCalled bool
}

func (s *HTTPNotePresenterSpy) SetResponseWriter(w http.ResponseWriter) {
	s.ResponseWriter = w
}

func (s *HTTPNotePresenterSpy) InternalServerError() {
	s.InternalServerErrorCalled = true
}

func NewHTTPPresenterSpy() *HTTPNotePresenterSpy {
	return &HTTPNotePresenterSpy{}
}
