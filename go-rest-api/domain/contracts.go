package domain

import "context"

// Application defines all operations implemented by the main application.
type Application interface {
	// CreateUser creates a new user and persists it into storage.
	CreateUser(ctx context.Context, in CreateUserInput) (string, []error)

	// // FindUser retrieves a specific user information from storage.
	// FindUser(ctx context.Context) []error

	// // ListUser retrieves several users information from storage.
	// ListUser(ctx context.Context) []error

	// // UpdateUser updates the information of an existing user.
	// UpdateUser(ctx context.Context) []error

	// // DeleteUser deletes an user.
	// DeleteUser(ctx context.Context) []error
}

// UsersRepository defines all operations performed by an user repository.
type UsersRepository interface {
	// Create creates and persists a new user into repository.
	Create(context.Context, User) error

	// // Find retrieves an specific user information from repository.
	// Find(ctx context.Context) []error

	// // List retrieves several users from repository.
	// List(ctx context.Context) []error

	// // Update updates the information of an existing user.
	// Update(ctx context.Context) []error

	// // SoftDelete flags user as deleted.
	// SoftDelete(ctx context.Context) []error

	// // Delete permanently deletes an user.
	// Delete(ctx context.Context) []error
}
