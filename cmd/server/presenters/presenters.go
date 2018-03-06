package presenters

import (
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type HTTPNotePresenterFactory interface {
	Create(http.ResponseWriter) HTTPNotePresenter
}

type HTTPNotePresenter interface {
	markdownnotes.CreatedNotePresenter
	markdownnotes.NotePresenter
	markdownnotes.NoteListPresenter
	ServiceUnavailable()
}
