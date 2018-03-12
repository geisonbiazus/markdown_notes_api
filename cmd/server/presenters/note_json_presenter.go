package presenters

import (
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type NoteJSONPresenter struct {
	JSONPresenter
}

func (p *NoteJSONPresenter) PresentCreatedNote(n markdownnotes.Note) {
	p.RenderJSON(http.StatusCreated, n)
}

func (p *NoteJSONPresenter) PresentUpdatedNote(n markdownnotes.Note) {
	p.PresentNote(n)
}

func (p *NoteJSONPresenter) PresentNote(n markdownnotes.Note) {
	p.RenderJSON(http.StatusOK, n)
}

func (p *NoteJSONPresenter) PresentNoteList(n []markdownnotes.Note) {
	p.RenderJSON(http.StatusOK, n)
}

type NoteJSONPresenterFactory struct {
	PresentNoteStatus int
}

func NewNoteJSONPresenterFactory() NotePresenterFactory {
	return NoteJSONPresenterFactory{}
}

func (f NoteJSONPresenterFactory) Create(w http.ResponseWriter) markdownnotes.NotePresenter {
	return &NoteJSONPresenter{JSONPresenter{w}}
}
