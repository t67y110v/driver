package handlers

import (
	"sync"

	"github.com/t67y110v/driver/internal/driver/config"
	"github.com/t67y110v/driver/internal/driver/logging"
	"github.com/t67y110v/driver/internal/driver/scale"
	pb "github.com/t67y110v/driver/pkg/driver"
)

type ScalesHandler struct {
	pb.ApiCallerScaleClient
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

// func newClient(c config.Config, ctx context.Context) (pb.ApiCallerScaleClient, error) {

// 	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", c.ScaleIP, c.ScalePort))
// 	if err != nil {
// 		return nil, err
// 	}

// 	client := pb.NewApiCallerScaleClient(conn)

// 	return client, nil
// }
