package note

import "github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"

type CreateNoteUseCase struct {
	NoteStorage markdownnotes.NoteStorage
}

func (u *CreateNoteUseCase) Run(title, content string) {
	newNote := markdownnotes.Note{
		Title:   title,
		Content: content,
	}

	u.NoteStorage.Save(newNote)
}
