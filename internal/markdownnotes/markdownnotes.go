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
	PresentCreatedNote(Note)
	PresentUpdatedNote(Note)
	PresentNote(Note)
	PresentNoteList(notes []Note)
	PresentValidationErrors([]ValidationError)
	PresentNotFound()
	PresentError(error)
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
