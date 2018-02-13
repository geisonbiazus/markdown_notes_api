package validators

import (
	"testing"

	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
	"github.com/stretchr/testify/assert"
)

func TestNoteValidator(t *testing.T) {
	validator := NoteValidator{}

	t.Run("It validates empty title", func(t *testing.T) {
		note := markdownnotes.Note{}
		expectedErrs := []markdownnotes.ValidationError{
			markdownnotes.ValidationError{
				Field:   "title",
				Message: "Can't be blank",
				Type:    "REQUIRED",
			},
		}

		errs := validator.Validate(note)

		assert.Equal(t, expectedErrs, errs)
	})
}
