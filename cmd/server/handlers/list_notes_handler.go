package handlers

import (
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/cmd/server/presenters"
	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type ListNotesHandler struct {
	UseCase          ListNoteUseCase
	PresenterFactory presenters.HTTPNotePresenterFactory
}

func NewListNotesHandler(u ListNoteUseCase, f presenters.HTTPNotePresenterFactory) *ListNotesHandler {
	return &ListNotesHandler{u, f}
}

func (h *ListNotesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	presenter := h.PresenterFactory.Create(w)
	err := h.UseCase.ListNotes(presenter)

	if err != nil {
		presenter.ServiceUnavailable()
	}
}

type ListNoteUseCase interface {
	ListNotes(markdownnotes.NoteListPresenter) error
}
