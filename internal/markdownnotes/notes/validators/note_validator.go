package validators

import (
	"errors"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

var WrongTypeError = errors.New("Type Error: It should be a markdownnotes.Note")

type NoteValidator struct{}

func (v NoteValidator) Validate(note interface{}) ([]markdownnotes.ValidationError, error) {
	n, ok := note.(markdownnotes.Note)
	if !ok {
		return []markdownnotes.ValidationError{}, WrongTypeError
	}

	errs := []markdownnotes.ValidationError{}

	if n.Title == "" {
		errs = append(errs, markdownnotes.ValidationError{
			Field: "title", Type: "REQUIRED", Message: "Can't be blank",
		})
	}
	return errs, nil
}
