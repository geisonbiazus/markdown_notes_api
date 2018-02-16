package main

import (
	"net/http"

	"github.com/geisonbiazus/markdown_notes_api/cmd/server/app"
)

func main() {
	router := app.InitRouter()
	http.ListenAndServe(":8080", router)
}
