package users

import (
	"strings"

	"github.com/pborman/uuid"
)

// UserID type
type UserID string

// User struct
type User struct {
	ID        UserID `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
}

// Users func
type Users struct {
	Users []User `json:"users,omitempty"`
}

// UserRepository interface get list user
type UserRepository interface {
	List() (*Users, error)
	Store(user *User) error
	Remove(userID UserID) error
	Update(user *User, id UserID) error
	Find(UserID UserID) (*User, error)
}

// NextUserID generates a new tracking ID.
func NextUserID() UserID {
	return UserID(strings.Split(strings.ToUpper(uuid.New()), "-")[0])
}
