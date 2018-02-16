package main

import (
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/cmd/server/handlers"
	"github.com/geisonbiazus/markdown_notes_api/cmd/server/presenters"
	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes/notes"
	"github.com/geisonbiazus/markdown_notes_api/internal/testing/fakes"
)

func main() {
	noteStorage := fakes.NewNoteStorageFake()
	createNoteUseCase := notes.NewCreateNoteUseCase(noteStorage)
	createNotePresenter := presenters.NewNoteJSONPresenter()
	createNoteHandler := handlers.NewCreateNoteHandler(createNoteUseCase, createNotePresenter)

	mux := http.NewServeMux()
	mux.Handle("/api/notes", createNoteHandler)

	http.ListenAndServe(":8080", mux)
}
