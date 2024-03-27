package store

import "github.com/T1mofey4/restAPI/internal/app/model"

// User repository
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
