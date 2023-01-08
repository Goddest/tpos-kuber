package store
type Store interface {
	User() UserRepository
	Notes() NoteRepository
}
