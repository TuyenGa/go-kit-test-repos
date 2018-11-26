package users

import (
	"context"
)

// Service user interface
type Service interface {
	RequestNewUser(ctx context.Context, user User) error
	List(ctx context.Context) (*Users, error)
	Update(ctx context.Context, user User, id UserID) error
	Find(ctx context.Context, id UserID) (*User, error)
	Delete(ctx context.Context, id UserID) error
}

type service struct {
	users UserRepository
}

func (s service) RequestNewUser(_ context.Context, user User) error {
	return nil
}

func (s service) List(_ context.Context) (*Users, error) {
	return &Users{}, nil
}

func (s service) Update(_ context.Context, user User, id UserID) error {
	return nil
}

func (s service) Find(_ context.Context, id UserID) (*User, error) {
	return &User{}, nil
}

func (s service) Delete(_ context.Context, id UserID) error {
	return nil
}

// NewService func make new Service struct userRepository
func NewService(users UserRepository) Service {
	return &service{
		users: users,
	}
}
