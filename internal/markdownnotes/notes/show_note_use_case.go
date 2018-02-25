package notes

import "github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"

type ShowNoteUseCase struct {
	NoteStorage markdownnotes.NoteStorage
}

func NewShowNoteUseCase(storage markdownnotes.NoteStorage) *ShowNoteUseCase {
	return &ShowNoteUseCase{storage}
}

func (u *ShowNoteUseCase) Run(noteID int, presenter markdownnotes.NotePresenter) error {
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
