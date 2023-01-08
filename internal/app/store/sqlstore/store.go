package sqlstore

import (
	"database/sql"

	"github.com/Goddest/tpos-kuber/internal/app/store"
	_ "github.com/lib/pq" // use pg driver
)
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
	noteRepository *NoteRepository
}
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}
func (s *Store) Notes() store.NoteRepository {
	if s.noteRepository != nil {
		return s.noteRepository
	}
	s.noteRepository = &NoteRepository{
		store: s,
	}
	return s.noteRepository
}
