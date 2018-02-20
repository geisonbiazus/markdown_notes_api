package presenters

import (
	"encoding/json"
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type NoteJSONPresenter struct {
	ResponseWriter http.ResponseWriter
}

func (p *NoteJSONPresenter) PresentNote(n markdownnotes.Note) {
	p.renderJSON(http.StatusCreated, n)
}

type errsContainer struct {
	Errors []markdownnotes.ValidationError `json:"errors"`
}

func (p *NoteJSONPresenter) PresentError(errs []markdownnotes.ValidationError) {
	p.renderJSON(http.StatusUnprocessableEntity, errsContainer{errs})
}

func (p *NoteJSONPresenter) ServiceUnavailable() {
	err := []markdownnotes.ValidationError{
		markdownnotes.ValidationError{
			Type: "SERVICE_UNAVAILABLE", Message: "Service Unavailable",
		},
	}
	p.renderJSON(http.StatusServiceUnavailable, errsContainer{err})
}

const contentType = "application/json"

func (p *NoteJSONPresenter) renderJSON(status int, object interface{}) {
	body, _ := json.Marshal(object)
	p.ResponseWriter.Header().Set("Content-Type", contentType)
	p.ResponseWriter.WriteHeader(status)
	p.ResponseWriter.Write(body)
}

type NoteJSONPresenterFactory struct{}

func (f NoteJSONPresenterFactory) Create(w http.ResponseWriter) HTTPNotePresenter {
	return &NoteJSONPresenter{w}
}

func NewNoteJSONPresenterFactory() HTTPNotePresenterFactory {
	return NoteJSONPresenterFactory{}
}
