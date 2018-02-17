package app

import (
	"net/http"
)

func InitRouter() http.Handler {
	handlers := InitHandlers()

	mux := http.NewServeMux()
	mux.Handle("/api/v1/notes", handlers.CreateNote)

	return mux
}
