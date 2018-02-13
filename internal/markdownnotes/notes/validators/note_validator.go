package validators

import (
	"errors"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

var WrongTypeError = errors.New("Type Error: It should be a markdownnotes.Note")

type NoteValidator struct{}

func (v NoteValidator) Validate(note markdownnotes.Note) []markdownnotes.ValidationError {
	errs := []markdownnotes.ValidationError{}

	if note.Title == "" {
		errs = append(errs, markdownnotes.ValidationError{
			Field: "title", Type: "REQUIRED", Message: "Can't be blank",
		})
	}
	return errs
}
