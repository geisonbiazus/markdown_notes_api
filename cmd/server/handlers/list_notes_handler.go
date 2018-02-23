package handlers

import (
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

type ListNotesHandler struct {
	UseCase          ListNoteUseCase
	PresenterFactory HTTPNoteListPresenterFactory
}

func NewListNotesHandler(u ListNoteUseCase, f HTTPNoteListPresenterFactory) *ListNotesHandler {
	return &ListNotesHandler{u, f}
}

func (h *ListNotesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	presenter := h.PresenterFactory.Create(w)
	err := h.UseCase.Run(presenter)

	if err != nil {
		presenter.ServiceUnavailable()
	}
}

type ListNoteUseCase interface {
	Run(markdownnotes.NoteListPresenter) error
}
type HTTPNoteListPresenter interface {
	markdownnotes.NoteListPresenter
	ServiceUnavailable()
}
type HTTPNoteListPresenterFactory interface {
	Create(w http.ResponseWriter) HTTPNoteListPresenter
}
