package sqlstore

import (
	"database/sql"

	"github.com/T1mofey4/restAPI/internal/app/store"
	_ "github.com/jackc/pgx/stdlib"
)

// Store
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// New
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
