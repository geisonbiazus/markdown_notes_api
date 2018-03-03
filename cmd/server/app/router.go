package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func InitRouter() http.Handler {
	h := InitHandlers()

	router := httprouter.New()
	router.Handler(http.MethodGet, "/api/v1/notes", h.ListNotes)
	router.Handler(http.MethodPost, "/api/v1/notes", h.CreateNote)
	router.Handler(http.MethodGet, "/api/v1/notes/:id", h.ShowNote)

	return router
}
