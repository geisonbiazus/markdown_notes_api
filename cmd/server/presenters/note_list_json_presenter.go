package presenters

import (
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type NoteListJSONPresenter struct {
	JSONPresenter
}

func (p *NoteListJSONPresenter) PresentNotes(n []markdownnotes.Note) {
	p.RenderJSON(http.StatusOK, n)
}

type NoteListJSONPresenterFactory struct{}

func (f NoteListJSONPresenterFactory) Create(w http.ResponseWriter) HTTPNoteListPresenter {
	return &NoteListJSONPresenter{JSONPresenter{w}}
}

func NewNoteListJSONPresenterFactory() HTTPNoteListPresenterFactory {
	return NoteListJSONPresenterFactory{}
}
