package main

import (
	"net/http"
	"os"

	"github.com/geisonbiazus/markdown_notes_api/cmd/server/app"
)

func main() {
	router := app.InitRouter()
	http.ListenAndServe(":"+getPort(), router)
}

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	return port
}
