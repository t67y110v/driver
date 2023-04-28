package handlers

import (
	"context"
	"io"
	"reflect"

	"github.com/t67y110v/driver/internal/driver/model"
	"github.com/t67y110v/driver/internal/driver/parser"
	"github.com/t67y110v/driver/internal/driver/scale"
	"github.com/t67y110v/driver/pkg/driver"
)

func (s *ScalesHandler) ScalesMessageOutChannel(stream driver.ApiCallerScale_ScalesMessageOutChannelServer) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for {

		stream.Send(&driver.ResponseScale{
			Message: string(s.scale.MakeMessage([]byte{scale.Commands["CMD_SET_ZERO"]})),
		})

		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		msg := []byte(req.Message)

		r := &model.SetZeroResponse{}
		err = r.Read(msg)
		if err != nil {
			continue
		}
		parser.FillResponseStruct(r.Raw(), reflect.ValueOf(&r).Elem())
		if r.Response.Raw().Meta.CommandCode == scale.Commands["CMD_ACK_SET"] {
			s.logger.Infoln("weight value set to zero")
		}

	}

}

func (s *ScalesHandler) GetInstantWeight(context.Context, *driver.Empty) (*driver.ResponseInstantWeight, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return nil, nil
}

func (s *ScalesHandler) GetState(ctx context.Context, empty *driver.Empty) (*driver.ResponseScale, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	message := s.scale.MakeMessage([]byte{scale.Commands["CMD_GET_MASSA"]})

	response := &driver.ResponseScale{
		Error:   scale.Errors[0x11],
		Message: string(message),
		Type:    "",
		Subtype: "",
	}
	return response, nil
}

func (s *ScalesHandler) SetTare(ctx context.Context, _ *driver.Empty) (*driver.ResponseSetScale, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// conn, err := newConnection(*s.config, ctx)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (s *ScalesHandler) SetTareValue(ctx context.Context, req *driver.RequestTareValue) (*driver.ResponseSetScale, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// message := s.scale.MakeMessagefForSetValue([]byte{scale.Commands["CMD_SET_ZERO"]}, []byte(req.Message))

	return nil, nil
}

func (s *ScalesHandler) SetZero(context.Context, *driver.Empty) (*driver.ResponseSetScale, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// message := s.scale.MakeMessage([]byte{scale.Commands["CMD_SET_ZERO"]})

	// response := &driver.ResponseSetScale{}
	// return response, nil

	return nil, nil
}
