package notes

import (
	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"
	"github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes/notes/validators"
)

type NoteUseCase struct {
	NoteStorage markdownnotes.NoteStorage
	Validator   validators.NoteValidator
}

func NewNoteUseCase(storage markdownnotes.NoteStorage) *NoteUseCase {
	return &NoteUseCase{
		NoteStorage: storage,
		Validator:   validators.NoteValidator{},
	}
}

func (u *NoteUseCase) CreateNote(title, content string, presenter markdownnotes.NotePresenter) error {
	note := markdownnotes.Note{
		Title:   title,
		Content: content,
	}

	errs := u.Validator.Validate(note)
	if len(errs) > 0 {
		presenter.PresentErrors(errs)
		return nil
	}

	note, err := u.NoteStorage.Save(note)
	if err != nil {
		return err
	}
	presenter.PresentNote(note)

	return nil
}

func (u *NoteUseCase) ListNotes(presenter markdownnotes.NotePresenter) error {
	notes, err := u.NoteStorage.FindAll()
	if err != nil {
		return err
	}
	presenter.PresentNotes(notes)
	return nil
}

func (u *NoteUseCase) ShowNote(noteID int, presenter markdownnotes.NotePresenter) error {
	note, err := u.NoteStorage.FindByID(noteID)
	if err != nil {
		return err
	}

	if note == markdownnotes.NoNote {
		presenter.NotFound()
		return nil
	}

	presenter.PresentNote(note)

	return nil
}

func (u *NoteUseCase) UpdateNote(noteID int, title, content string, presenter markdownnotes.NotePresenter) error {
	note, err := u.NoteStorage.FindByID(noteID)
	if err != nil {
		return err
	}

	if note == markdownnotes.NoNote {
		presenter.NotFound()
		return nil
	}

	note.Title = title
	note.Content = content

	errs := u.Validator.Validate(note)
	if len(errs) > 0 {
		presenter.PresentErrors(errs)
		return nil
	}

	note, err = u.NoteStorage.Save(note)
	if err != nil {
		return err
	}

	presenter.PresentNote(note)

	return nil
}
