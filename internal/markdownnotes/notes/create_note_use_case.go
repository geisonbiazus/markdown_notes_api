package notes

import "github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"

type CreateNoteUseCase struct {
	NoteStorage   markdownnotes.NoteStorage
	NotePresenter NotePresenter
}

func (u *CreateNoteUseCase) Run(title, content string) error {
	note := markdownnotes.Note{
		Title:   title,
		Content: content,
	}

	errs := validateNote(note)

	if len(errs) > 0 {
		u.NotePresenter.PresentError(errs)
		return nil
	}

	note, err := u.NoteStorage.Save(note)
	u.NotePresenter.PresentNote(note)

	return err
}

func validateNote(note markdownnotes.Note) []markdownnotes.ValidationError {
	errs := []markdownnotes.ValidationError{}

	if note.Title == "" {
		errs = append(errs, markdownnotes.ValidationError{
			Field: "title", Type: "REQUIRED", Message: "Can't be blank",
		})
	}
	return errs
}

type NotePresenter interface {
	PresentNote(markdownnotes.Note)
	PresentError([]markdownnotes.ValidationError)
}
