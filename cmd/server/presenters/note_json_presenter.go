package presenters

import (
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type NoteJSONPresenter struct {
	PresentNoteStatus int
	JSONPresenter
}

func (p *NoteJSONPresenter) PresentNote(n markdownnotes.Note) {
	p.RenderJSON(p.PresentNoteStatus, n)
}

type NoteJSONPresenterFactory struct {
	PresentNoteStatus int
}

func (f NoteJSONPresenterFactory) Create(w http.ResponseWriter) HTTPNotePresenter {
	return &NoteJSONPresenter{f.PresentNoteStatus, JSONPresenter{w}}
}

func NewNoteJSONPresenterFactory(presentNoteStatus int) HTTPNotePresenterFactory {
	return NoteJSONPresenterFactory{PresentNoteStatus: presentNoteStatus}
}
