package fakes

import "github.com/geisonbiazus/markdown_notes_api/internal/markdownnotes"

type NoteStorageFake struct {
	data   map[int]markdownnotes.Note
	lastID int
}

func (f *NoteStorageFake) Save(n markdownnotes.Note) (markdownnotes.Note, error) {
	if n.ID == 0 {
		f.lastID++
		n.ID = f.lastID
	}
	f.data[n.ID] = n
	return n, nil
}

func (f *NoteStorageFake) FindAll() ([]markdownnotes.Note, error) {
	notes := []markdownnotes.Note{}
	for _, n := range f.data {
		notes = append(notes, n)
	}
	return notes, nil
}

func (f *NoteStorageFake) FindByID(id int) (markdownnotes.Note, error) {
	for _, note := range f.data {
		if note.ID == id {
			return note, nil
		}
	}
	return markdownnotes.NoNote, nil
}

func NewNoteStorageFake() *NoteStorageFake {
	return &NoteStorageFake{data: make(map[int]markdownnotes.Note)}
}
