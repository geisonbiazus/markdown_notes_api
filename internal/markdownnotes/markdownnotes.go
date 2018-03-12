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
	PresentNotes(notes []Note)
	PresentErrors([]ValidationError)
	NotFound()
	ServiceUnavailable()
}

type NoteUseCase interface {
	CreateNote(title, content string, p NotePresenter) error
	ShowNote(id int, p NotePresenter) error
	UpdateNote(id int, title, content string, p NotePresenter) error
	ListNotes(p NotePresenter) error
}

type ValidationError struct {
	Field   string `json:"field"`
	Type    string `json:"type"`
	Message string `json:"message"`
}
