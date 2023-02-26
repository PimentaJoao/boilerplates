package app

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pimentajoao/boilerplates/go-rest-api/domain"
)

func (app application) CreateUser(ctx context.Context, in domain.CreateUserInput) (string, []error) {
	// err := validator.ValidateProduct(in)
	// if err != nil {
	// 	return "", []error{err}
	// }

	user := domain.User{
		ExternalID:  uuid.New().String(),
		Name:        in.Name,
		Description: in.Description,
		Collection:  in.Collection,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := app.usersRepo.Create(ctx, user)
	if err != nil {
		return "", []error{err}
	}

	return user.ExternalID, nil
}
