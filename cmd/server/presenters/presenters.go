package presenters

import (
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type NotePresenterFactory interface {
	Create(http.ResponseWriter) markdownnotes.NotePresenter
}
