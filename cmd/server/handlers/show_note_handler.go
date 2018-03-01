package handlers

import (
	"net/http"
	"strconv"

	"github.com/geisonbiazus/markdown_notes_api/cmd/server/presenters"
	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
	"github.com/julienschmidt/httprouter"
)

type ShowNoteHandler struct {
	UseCase          ShowNoteUseCase
	PresenterFactory presenters.HTTPNotePresenterFactory
}

func NewShowNoteHandler(u ShowNoteUseCase, f presenters.HTTPNotePresenterFactory) *ShowNoteHandler {
	return &ShowNoteHandler{UseCase: u, PresenterFactory: f}
}

func (h *ShowNoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	noteID, _ := strconv.Atoi(p.ByName("id"))
	presenter := h.PresenterFactory.Create(w)
	err := h.UseCase.Run(noteID, presenter)
	if err != nil {
		presenter.ServiceUnavailable()
	}
}

type ShowNoteUseCase interface {
	Run(noteID int, p markdownnotes.NotePresenter) error
}
