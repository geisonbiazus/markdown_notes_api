package handlers

import (
	"net/http"
	"strconv"

	"github.com/geisonbiazus/markdown_notes_api/cmd/server/presenters"
	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
	"github.com/julienschmidt/httprouter"
)

type ShowNoteHandler struct {
	UseCase          markdownnotes.NoteUseCase
	PresenterFactory presenters.NotePresenterFactory
}

func NewShowNoteHandler(u markdownnotes.NoteUseCase, f presenters.NotePresenterFactory) *ShowNoteHandler {
	return &ShowNoteHandler{UseCase: u, PresenterFactory: f}
}

func (h *ShowNoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := httprouter.ParamsFromContext(r.Context())
	noteID, _ := strconv.Atoi(p.ByName("id"))
	presenter := h.PresenterFactory.Create(w)
	err := h.UseCase.ShowNote(noteID, presenter)
	if err != nil {
		presenter.PresentError(err)
	}
}

type ShowNoteUseCase interface {
	ShowNote(noteID int, p markdownnotes.NotePresenter) error
}
