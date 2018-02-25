package presenters

import (
	"encoding/json"
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type JSONPresenter struct {
	ResponseWriter http.ResponseWriter
}

const contentType = "application/json"

func (p *JSONPresenter) RenderJSON(status int, object interface{}) {
	body, _ := json.Marshal(object)
	p.ResponseWriter.Header().Set("Content-Type", contentType)
	p.ResponseWriter.WriteHeader(status)
	p.ResponseWriter.Write(body)
}

type errsContainer struct {
	Errors []markdownnotes.ValidationError `json:"errors"`
}

func (p *JSONPresenter) PresentError(errs []markdownnotes.ValidationError) {
	p.RenderJSON(http.StatusUnprocessableEntity, errsContainer{errs})
}

func (p *JSONPresenter) ServiceUnavailable() {
	err := []markdownnotes.ValidationError{
		markdownnotes.ValidationError{
			Type: "SERVICE_UNAVAILABLE", Message: "Service Unavailable",
		},
	}
	p.RenderJSON(http.StatusServiceUnavailable, errsContainer{err})
}

func (p *JSONPresenter) NotFound() {

}
