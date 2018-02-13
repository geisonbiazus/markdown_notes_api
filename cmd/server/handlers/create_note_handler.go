package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes/notes"
)

type createNoteHandlerParams struct {
	Note struct {
		Title   string
		Content string
	}
}

type CreateNoteHandler struct {
	UseCase   CreateNoteUseCase
	Presenter HTTPNotePresenter
}

func (h *CreateNoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Presenter.SetResponseWriter(w)
	params := h.getParams(r)
	h.UseCase.Run(params.Note.Title, params.Note.Content, h.Presenter)
}

func (h *CreateNoteHandler) getParams(r *http.Request) createNoteHandlerParams {
	params := createNoteHandlerParams{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &params)
	return params
}

func NewCreateNoteHandler(u CreateNoteUseCase, p HTTPNotePresenter) *CreateNoteHandler {
	return &CreateNoteHandler{UseCase: u, Presenter: p}
}

type CreateNoteUseCase interface {
	Run(title, content string, presenter notes.NotePresenter) error
}

type HTTPNotePresenter interface {
	notes.NotePresenter
	SetResponseWriter(http.ResponseWriter)
}
