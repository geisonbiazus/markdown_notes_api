package doubles

import (
	"net/http"
)

type HTTPNotePresenterSpy struct {
	*NotePresenterSpy
	ResponseWriter http.ResponseWriter
}

func (s *HTTPNotePresenterSpy) SetResponseWriter(w http.ResponseWriter) {
	s.ResponseWriter = w
}

func NewHTTPPresenterSpy() *HTTPNotePresenterSpy {
	return &HTTPNotePresenterSpy{}
}
