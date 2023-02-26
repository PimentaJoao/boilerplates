package main

import (
	"context"
	"os"
	"strconv"

	repository "github.com/pimentajoao/boilerplates/go-rest-api/adapter"
	"github.com/pimentajoao/boilerplates/go-rest-api/app"
	"github.com/pimentajoao/boilerplates/go-rest-api/ports/rest"
)

func main() {
	ctx := context.Background()

	usersRepo := repository.MustNewUsersRepository(repository.UsersRepositoryParams{
		DBConn: nil,
	})

	app := app.MustNewApplication(app.ApplicationParams{
		UsersRepo: usersRepo,
	})

	port, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		panic("Missing environment variable: API_PORT")
	}
	server := rest.MustNewServer(rest.ServerParams{
		Application: app,
		Port:        port,
	})

	server.Run(ctx)
}
