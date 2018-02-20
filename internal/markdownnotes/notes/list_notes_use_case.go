package notes

import "github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"

type ListNotesUseCase struct {
	NoteStorage markdownnotes.NoteStorage
}

func NewListNotesUseCase(storage markdownnotes.NoteStorage) *ListNotesUseCase {
	return &ListNotesUseCase{storage}
}

func (u *ListNotesUseCase) Run(presenter markdownnotes.NoteListPresenter) error {
	notes, err := u.NoteStorage.FindAll()
	if err != nil {
		return err
	}
	presenter.PresentNotes(notes)
	return nil
}
