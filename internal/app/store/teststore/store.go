package teststore

import (
	"github.com/Goddest/tpos-kuber/internal/app/model"
	"github.com/Goddest/tpos-kuber/internal/app/store"
)
type Store struct {
	userRepository *UserRepository
	noteRepository *NoteRepository
}
func New() *Store {
	return &Store{}
}
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}
	return s.userRepository
}
func (s *Store) Notes() store.NoteRepository {
	if s.noteRepository != nil {
		return s.noteRepository
	}
	s.noteRepository = &NoteRepository{
		store: s,
		notes: make(map[int]*model.Note),
	}
	return s.noteRepository
}
