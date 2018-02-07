package notes

import "github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"

type CreateNoteUseCase struct {
	NoteStorage markdownnotes.NoteStorage
}

func (u *CreateNoteUseCase) Run(title, content string) error {
	note := markdownnotes.Note{
		Title:   title,
		Content: content,
	}

	return u.NoteStorage.Save(note)
}
