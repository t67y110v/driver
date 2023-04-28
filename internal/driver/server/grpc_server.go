package server

import (
	"context"
	"fmt"
	"net"

	"github.com/t67y110v/driver/internal/driver/config"
	"github.com/t67y110v/driver/internal/driver/logging"
	"github.com/t67y110v/driver/pkg/driver"
	"google.golang.org/grpc"
)

func Start(config *config.Config) error {

	ctx := context.Background()

	logger := logging.GetLogger()

	server := newServer(&logger, config)

	return server.ListenAndServe(ctx)
}

func (s *Server) ListenAndServe(ctx context.Context) error {
	serv := grpc.NewServer()

	driver.RegisterApiCallerScaleServer(serv, s.handlres)

	lisAddr := fmt.Sprintf("0.0.0.0:%s", s.config.ServerPort)

	lis, err := net.Listen("tcp", lisAddr)
	if err != nil {
		return err
	}

	return serv.Serve(lis)

}
