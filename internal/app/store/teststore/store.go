package teststore

import (
	"github.com/T1mofey4/restAPI/internal/app/model"
	"github.com/T1mofey4/restAPI/internal/app/store"
	_ "github.com/jackc/pgx/stdlib"
)

// Store
type Store struct {
	userRepository *UserRepository
}

// New
func New() *Store {
	return &Store{}
}

// User
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
