package markdownnotes

type Note struct {
	Title   string
	Content string
}

type NoteStorage interface {
	Save(note Note) error
}
