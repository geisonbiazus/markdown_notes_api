package app

import (
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/cmd/server/handlers"
	"github.com/geisonbiazus/markdown_notes_api/cmd/server/presenters"
	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes/notes"
	"github.com/geisonbiazus/markdown_notes_api/internal/testing/fakes"
)

type Handlers struct {
	CreateNote http.Handler
}

func InitHandlers() *Handlers {
	noteStorage := fakes.NewNoteStorageFake()
	createNoteUseCase := notes.NewCreateNoteUseCase(noteStorage)
	createNotePresenter := presenters.NewNoteJSONPresenter()

	return &Handlers{
		CreateNote: handlers.NewCreateNoteHandler(createNoteUseCase, createNotePresenter),
	}
}
