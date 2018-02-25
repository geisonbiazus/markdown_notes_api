package presenters

import (
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type NoteJSONPresenter struct {
	JSONPresenter
}

func (p *NoteJSONPresenter) PresentNote(n markdownnotes.Note) {
	p.RenderJSON(http.StatusCreated, n)
}

type NoteJSONPresenterFactory struct{}

func (f NoteJSONPresenterFactory) Create(w http.ResponseWriter) HTTPNotePresenter {
	return &NoteJSONPresenter{JSONPresenter{w}}
}

func NewNoteJSONPresenterFactory() HTTPNotePresenterFactory {
	return NoteJSONPresenterFactory{}
}
