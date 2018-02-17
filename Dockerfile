FROM golang:1.9.4

COPY . $GOPATH/src/github.com/geisonbiazus/markdown_notes_api
WORKDIR $GOPATH/src/github.com/geisonbiazus/markdown_notes_api

RUN make build

CMD markdown_notes_api

EXPOSE 8080
