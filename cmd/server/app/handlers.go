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
	ListNotes  http.Handler
}

func InitHandlers() *Handlers {
	noteStorage := fakes.NewNoteStorageFake()
	createNoteUseCase := notes.NewCreateNoteUseCase(noteStorage)
	listNotesUseCase := notes.NewListNotesUseCase(noteStorage)

	createNotePresenterFactory := presenters.NewNoteJSONPresenterFactory()
	noteListPresenterFactory := presenters.NewNoteListJSONPresenterFactory()

	return &Handlers{
		CreateNote: handlers.NewCreateNoteHandler(createNoteUseCase, createNotePresenterFactory),
		ListNotes:  handlers.NewListNotesHandler(listNotesUseCase, noteListPresenterFactory),
	}
}
