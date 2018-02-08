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

	note, err := u.NoteStorage.Save(note)
	u.NotePresenter.PresentNote(note)

	return err
}

type NotePresenter interface {
	PresentNote(markdownnotes.Note)
}
