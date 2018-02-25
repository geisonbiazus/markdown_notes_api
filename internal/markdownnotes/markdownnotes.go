package markdownnotes

var NoNote = Note{}

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NoteStorage interface {
	Save(note Note) (Note, error)
	FindAll() ([]Note, error)
	FindByID(id int) (Note, error)
}

type NotePresenter interface {
	PresentNote(Note)
	PresentError([]ValidationError)
	NotFound()
}

type NoteListPresenter interface {
	PresentNotes(notes []Note)
}

type ValidationError struct {
	Field   string `json:"field"`
	Type    string `json:"type"`
	Message string `json:"message"`
}
