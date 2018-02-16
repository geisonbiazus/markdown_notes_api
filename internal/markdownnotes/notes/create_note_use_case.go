package notes

import (
	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes/notes/validators"
)

type CreateNoteUseCase struct {
	NoteStorage markdownnotes.NoteStorage
	Validator   Validator
}

func (u *CreateNoteUseCase) Run(title, content string, presenter markdownnotes.NotePresenter) error {
	note := markdownnotes.Note{
		Title:   title,
		Content: content,
	}

	errs := u.Validator.Validate(note)
	if len(errs) > 0 {
		presenter.PresentError(errs)
		return nil
	}

	note, err := u.NoteStorage.Save(note)
	if err != nil {
		return err
	}
	presenter.PresentNote(note)

	return nil
}

func NewCreateNoteUseCase(storage markdownnotes.NoteStorage) *CreateNoteUseCase {
	return &CreateNoteUseCase{
		NoteStorage: storage,
		Validator:   validators.NoteValidator{},
	}
}

type Validator interface {
	Validate(markdownnotes.Note) []markdownnotes.ValidationError
}
