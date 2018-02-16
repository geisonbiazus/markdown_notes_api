package doubles

import (
	"net/http"
)

type HTTPNotePresenterSpy struct {
	*NotePresenterSpy
	ResponseWriter           http.ResponseWriter
	ServiceUnavailableCalled bool
}

func (s *HTTPNotePresenterSpy) SetResponseWriter(w http.ResponseWriter) {
	s.ResponseWriter = w
}

func (s *HTTPNotePresenterSpy) ServiceUnavailable() {
	s.ServiceUnavailableCalled = true
}

func NewHTTPPresenterSpy() *HTTPNotePresenterSpy {
	return &HTTPNotePresenterSpy{}
}
