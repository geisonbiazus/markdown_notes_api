package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/geisonbiazus/markdown_notes_api/cmd/server/presenters"
	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
	"github.com/julienschmidt/httprouter"
)

type UpdateNoteHandler struct {
	UseCase          markdownnotes.NoteUseCase
	PresenterFactory presenters.NotePresenterFactory
}

func NewUpdateNoteHandler(u markdownnotes.NoteUseCase, f presenters.NotePresenterFactory) *UpdateNoteHandler {
	return &UpdateNoteHandler{UseCase: u, PresenterFactory: f}
}

func (h *UpdateNoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, body := h.getParams(r)
	presenter := h.PresenterFactory.Create(w)
	err := h.UseCase.UpdateNote(id, body.Note.Title, body.Note.Content, presenter)

	if err != nil {
		presenter.PresentError(err)
	}
}

func (h *UpdateNoteHandler) getParams(r *http.Request) (int, noteHandlerParams) {
	params := httprouter.ParamsFromContext(r.Context())
	id, _ := strconv.Atoi(params.ByName("id"))
	body := noteHandlerParams{}
	json.NewDecoder(r.Body).Decode(&body)
	return id, body
}
