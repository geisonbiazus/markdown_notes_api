package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/cmd/server/presenters"
	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type createNoteHandlerParams struct {
	Note struct {
		Title   string
		Content string
	}
}

type CreateNoteHandler struct {
	UseCase          markdownnotes.NoteUseCase
	PresenterFactory presenters.NotePresenterFactory
}

func NewCreateNoteHandler(u markdownnotes.NoteUseCase, pf presenters.NotePresenterFactory) *CreateNoteHandler {
	return &CreateNoteHandler{UseCase: u, PresenterFactory: pf}
}

func (h *CreateNoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	presenter := h.PresenterFactory.Create(w)
	params := h.getParams(r)

	err := h.UseCase.CreateNote(params.Note.Title, params.Note.Content, presenter)
	if err != nil {
		presenter.PresentError(err)
	}
}

func (h *CreateNoteHandler) getParams(r *http.Request) createNoteHandlerParams {
	params := createNoteHandlerParams{}
	json.NewDecoder(r.Body).Decode(&params)
	return params
}
