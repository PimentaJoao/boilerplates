package domain

import (
	"time"

	"github.com/google/uuid"
)

// User defines all fields for an user.
type User struct {
	ExternalID  string
	Name        string
	Description string
	Collection  string

	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser takes an input and creates a complete product ready to be persisted in storage.
func NewUser(input CreateUserInput) User {
	return User{
		ExternalID:  uuid.New().String(),
		Name:        input.Name,
		Description: input.Description,
		Collection:  input.Collection,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
