package handlers

import (
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/cmd/server/presenters"
	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type ListNotesHandler struct {
	UseCase          markdownnotes.NoteUseCase
	PresenterFactory presenters.NotePresenterFactory
}

func NewListNotesHandler(u markdownnotes.NoteUseCase, f presenters.NotePresenterFactory) *ListNotesHandler {
	return &ListNotesHandler{u, f}
}

func (h *ListNotesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	presenter := h.PresenterFactory.Create(w)
	err := h.UseCase.ListNotes(presenter)

	if err != nil {
		presenter.ServiceUnavailable()
	}
}
