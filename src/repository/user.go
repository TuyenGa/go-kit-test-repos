package repository

import (
	"database/sql"

	"github.com/tuyenga/go-kit-test-repos/src/users"
)

type userRepository struct {
	db *sql.DB
}

func (repos *userRepository) List() (*users.Users, error) {
	return users.Users{}, nil
}

func (repos *userRepository) Store(user *users.User) error {
	return nil
}

func (repos *userRepository) Remove(id users.UserID) error {
	return nil
}

func (repos *userRepository) Update(user *users.User, id users.UserID) error {
	return nil
}

func (repos *userRepository) Find(id users.UserID) (*users.User, error) {
	return &users.User{}, nil
}

// NewRepository func make a new Repository db exec
func NewRepository(db *sql.DB) users.UserRepository {
	return &userRepository{
		db: db,
	}
}
