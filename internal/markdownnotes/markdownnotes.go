package markdownnotes

type Note struct {
	ID      int
	Title   string
	Content string
}

type NoteStorage interface {
	Save(note Note) (Note, error)
}

type Validator interface {
	Validate(interface{}) ([]ValidationError, error)
}

type ValidationError struct {
	Field   string
	Type    string
	Message string
}
