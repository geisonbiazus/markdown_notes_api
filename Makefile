run:
	go run cmd/server/main.go

test:
	go test ./cmd/... ./internal/...

build:
	go clean
	go build -o `go env GOPATH`/bin/markdown_notes_api ./cmd/server
