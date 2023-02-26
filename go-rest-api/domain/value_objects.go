package domain

// CreateUserInput defines all fields that are filled required in order to create a new user.
type CreateUserInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Collection  string `json:"collection"`
}
