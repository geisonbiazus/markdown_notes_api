package presenters

import (
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type HTTPNotePresenterFactory interface {
	Create(http.ResponseWriter) HTTPNotePresenter
}

type HTTPNotePresenter interface {
	markdownnotes.NotePresenter
	ServiceUnavailable()
}

type HTTPNoteListPresenter interface {
	markdownnotes.NoteListPresenter
	ServiceUnavailable()
}

type HTTPNoteListPresenterFactory interface {
	Create(w http.ResponseWriter) HTTPNoteListPresenter
}
