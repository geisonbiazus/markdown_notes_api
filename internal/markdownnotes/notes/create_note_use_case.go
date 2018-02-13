package notes

import (
	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes/notes/validators"
)

type CreateNoteUseCase struct {
	NoteStorage   markdownnotes.NoteStorage
	NotePresenter NotePresenter
	Validator     Validator
}

func (u *CreateNoteUseCase) Run(title, content string) error {
	note := markdownnotes.Note{
		Title:   title,
		Content: content,
	}

	errs := u.Validator.Validate(note)
	if len(errs) > 0 {
		u.NotePresenter.PresentError(errs)
		return nil
	}

	note, err := u.NoteStorage.Save(note)
	if err != nil {
		return err
	}
	u.NotePresenter.PresentNote(note)

	return nil
}

type CreateNoteUseCaseFactory struct {
	NoteStorage markdownnotes.NoteStorage
}

func (f *CreateNoteUseCaseFactory) Create(p NotePresenter) *CreateNoteUseCase {
	return &CreateNoteUseCase{
		NoteStorage:   f.NoteStorage,
		NotePresenter: p,
		Validator:     validators.NoteValidator{},
	}
}

type Validator interface {
	Validate(markdownnotes.Note) []markdownnotes.ValidationError
}

type NotePresenter interface {
	PresentNote(markdownnotes.Note)
	PresentError([]markdownnotes.ValidationError)
}
