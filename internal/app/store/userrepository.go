package store

import "github.com/T1mofey4/restAPI/internal/app/model"

// User repository
type UserRepository struct {
	store *Store
}

// Create
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

// Find by email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}
