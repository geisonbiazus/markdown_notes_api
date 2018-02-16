package validators

import "github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"

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
