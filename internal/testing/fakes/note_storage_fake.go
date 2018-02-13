package fakes

import "github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"

type NoteStorageFake struct {
	data   []markdownnotes.Note
	lastID int
}

func (f *NoteStorageFake) Save(n markdownnotes.Note) (markdownnotes.Note, error) {
	f.lastID++
	n.ID = f.lastID
	f.data = append(f.data, n)
	return n, nil
}
