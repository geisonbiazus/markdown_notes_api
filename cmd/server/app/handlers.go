package app

import (
	"github.com/geisonbiazus/markdown_notes_api/cmd/server/handlers"
	"github.com/geisonbiazus/markdown_notes_api/cmd/server/presenters"
	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes/notes"
	"github.com/geisonbiazus/markdown_notes_api/internal/testing/fakes"
)

type Handlers struct {
	CreateNote *handlers.CreateNoteHandler
	ListNotes  *handlers.ListNotesHandler
	ShowNote   *handlers.ShowNoteHandler
	UpdateNote *handlers.UpdateNoteHandler
}

func InitHandlers() *Handlers {
	noteStorage := fakes.NewNoteStorageFake()
	noteUseCase := notes.NewNoteUseCase(noteStorage)
	notePresenterFactory := presenters.NewNoteJSONPresenterFactory()

	return &Handlers{
		CreateNote: handlers.NewCreateNoteHandler(noteUseCase, notePresenterFactory),
		ListNotes:  handlers.NewListNotesHandler(noteUseCase, notePresenterFactory),
		ShowNote:   handlers.NewShowNoteHandler(noteUseCase, notePresenterFactory),
		UpdateNote: handlers.NewUpdateNoteHandler(noteUseCase, notePresenterFactory),
	}
}
