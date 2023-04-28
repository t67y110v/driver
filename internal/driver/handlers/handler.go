package handlers

import (
	"context"
	"fmt"
	"sync"

	"github.com/t67y110v/driver/internal/driver/config"
	"github.com/t67y110v/driver/internal/driver/logging"
	"github.com/t67y110v/driver/internal/driver/scale"
	"github.com/t67y110v/driver/pkg/driver"
	"google.golang.org/grpc"
)

type ScalesHandler struct {
	driver.UnimplementedApiCallerScaleServer
	logger *logging.Logger
	config *config.Config
	scale  *scale.Scale
	mu     *sync.RWMutex
}

func New(c *config.Config, l *logging.Logger) *ScalesHandler {

	h := &ScalesHandler{
		logger: l,
		config: c,
		scale:  scale.NewScale(),
		mu:     &sync.RWMutex{},
	}

	return h
}

func newConnection(c config.Config, ctx context.Context) (*grpc.ClientConn, error) {
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%s", c.ScaleIP, c.ScalePort))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
