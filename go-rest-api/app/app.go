package app

import (
	"fmt"

	"github.com/pimentajoao/boilerplates/go-rest-api/domain"
)

// ApplicationParams encapsulates all dependencies for an application.
type ApplicationParams struct {
	UsersRepo domain.UsersRepository
}

type application struct {
	usersRepo domain.UsersRepository
}

// NewApplication creates a new application, verifying if all dependencies are correctly provided.
func NewApplication(params ApplicationParams) (application, error) {
	if params.UsersRepo == nil {
		return application{}, fmt.Errorf("[application] missing dependency: UsersRepo")
	}

	return application{
		usersRepo: params.UsersRepo,
	}, nil
}

// MustNewApplication ensures that the application is created with all necessary dependencies,
// panicking if it's not the case.
func MustNewApplication(params ApplicationParams) application {
	app, err := NewApplication(params)
	if err != nil {
		panic(err)
	}
	return app
}
