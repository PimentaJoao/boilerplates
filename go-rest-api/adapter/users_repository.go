package repository

import (
	"context"
	"fmt"

	"github.com/pimentajoao/boilerplates/go-rest-api/domain"

	"gorm.io/gorm"
)

// UsersRepositoryParams encapsulates all dependencies for a Postgres users repository.
type UsersRepositoryParams struct {
	DBConn *gorm.DB
}

// UsersRepository defines a Postgres user repository.
type UsersRepository struct {
	dbConn *gorm.DB
}

// NewUsersRepository creates a new Postgres users repository, verifying if all
// dependencies are correctly provided.
func NewUsersRepository(params UsersRepositoryParams) (UsersRepository, error) {
	// if params.DBConn == nil {
	// 	return UsersRepository{}, fmt.Errorf("[application] missing dependency: DBConn")
	// }

	return UsersRepository{
		dbConn: params.DBConn,
	}, nil
}

// MustNewUsersRepository ensures that a Portgres user repository is created with
// all necessary dependencies, panicking if it's not the case.
func MustNewUsersRepository(repoParams UsersRepositoryParams) UsersRepository {
	repo, err := NewUsersRepository(repoParams)
	if err != nil {
		panic(err)
	}
	return repo
}

// Create creates a new item and persists it into storage.
func (upr UsersRepository) Create(ctx context.Context, u domain.User) error {
	return fmt.Errorf("[repository] error: not implemented yet")
}
