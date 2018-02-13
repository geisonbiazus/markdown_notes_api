package validators

import (
	"reflect"
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
)

func TestNoteValidator(t *testing.T) {
	validator := NoteValidator{}

	t.Run("Given an empty title, it returns a validation error", func(t *testing.T) {
		note := markdownnotes.Note{}
		expectedErrs := []markdownnotes.ValidationError{
			markdownnotes.ValidationError{
				Field:   "title",
				Message: "Can't be blank",
				Type:    "REQUIRED",
			},
		}

		errs := validator.Validate(note)

		if !reflect.DeepEqual(expectedErrs, errs) {
			t.Errorf("Expected: %v, Actual: %v", expectedErrs, errs)
		}
	})
}
