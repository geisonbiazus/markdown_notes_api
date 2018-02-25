package app

import (
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/cmd/server/handlers"
)

func InitRouter() http.Handler {
	h := InitHandlers()

	mux := http.NewServeMux()
	mux.Handle("/api/v1/notes", &handlers.HTTPMethodHandler{
		Post: h.CreateNote,
		Get:  h.ListNotes,
	})

	return mux
}
