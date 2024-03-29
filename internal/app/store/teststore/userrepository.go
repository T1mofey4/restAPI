package teststore

import (
	"errors"

	"github.com/T1mofey4/restAPI/internal/app/model"
)

// UserRepository
type UserRepository struct {
	store *Store
	users map[string]*model.User
}

// Create
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.Email] = u
	u.ID = len(r.users)

	return nil
}

// Find by email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, errors.New("User not found")
	}

	return u, nil
}
