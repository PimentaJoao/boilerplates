package rest

import (
	"context"
	"fmt"

	"github.com/pimentajoao/boilerplates/go-rest-api/domain"

	"github.com/gin-gonic/gin"
)

// ServerParams encapsulates all dependencies for a server.
type ServerParams struct {
	Port        int
	Application domain.Application
}

type server struct {
	port int
	app  domain.Application
}

// NewServer creates a new server, verifying if all dependencies are correctly provided.
func NewServer(params ServerParams) (server, error) {
	if params.Application == nil {
		return server{}, fmt.Errorf("[server] missing dependency: Application")
	}

	if params.Port == 0 {
		params.Port = 8080
	}

	return server{
		app:  params.Application,
		port: params.Port,
	}, nil
}

// MustNewServer ensures that the server is created with all necessary dependencies,
// panicking if it's not the case.
func MustNewServer(params ServerParams) server {
	server, err := NewServer(params)
	if err != nil {
		panic(err)
	}

	return server
}

// Run spins up the server.
func (s server) Run(ctx context.Context) {
	router := gin.Default()

	router.GET("/ping", s.PingEndpoint)

	port := fmt.Sprintf(":%d", s.port)
	router.Run(port)
}
