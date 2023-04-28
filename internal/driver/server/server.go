package server

import (
	"github.com/t67y110v/driver/internal/driver/config"
	"github.com/t67y110v/driver/internal/driver/handlers"
	"github.com/t67y110v/driver/internal/driver/logging"
)

type Server struct {
	logger   *logging.Logger
	config   *config.Config
	handlres *handlers.ScalesHandler
}

func newServer(l *logging.Logger, c *config.Config) *Server {
	server := &Server{
		logger:   l,
		config:   c,
		handlres: handlers.New(c, l),
	}

	return server
}
