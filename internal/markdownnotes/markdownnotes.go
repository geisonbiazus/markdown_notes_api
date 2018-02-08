package markdownnotes

type Note struct {
	ID      int
	Title   string
	Content string
}

type NoteStorage interface {
	Save(note Note) (Note, error)
}
