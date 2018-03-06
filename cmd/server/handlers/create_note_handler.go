package handlers

import (
	"encoding/json"
	"io/ioutil"
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
	UseCase          CreateNoteUseCase
	PresenterFactory presenters.HTTPNotePresenterFactory
}

func NewCreateNoteHandler(u CreateNoteUseCase, pf presenters.HTTPNotePresenterFactory) *CreateNoteHandler {
	return &CreateNoteHandler{UseCase: u, PresenterFactory: pf}
}

func (h *CreateNoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	presenter := h.PresenterFactory.Create(w)
	params := h.getParams(r)

	err := h.UseCase.CreateNote(params.Note.Title, params.Note.Content, presenter)
	if err != nil {
		presenter.ServiceUnavailable()
	}
}

func (h *CreateNoteHandler) getParams(r *http.Request) createNoteHandlerParams {
	params := createNoteHandlerParams{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &params)
	return params
}

type CreateNoteUseCase interface {
	CreateNote(title, content string, presenter markdownnotes.CreatedNotePresenter) error
}
