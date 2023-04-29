package handlers

import (
	"context"
	"fmt"
	"io"
	"time"

	// "io"
	// "reflect"

	// "github.com/t67y110v/driver/internal/driver/model"
	// "github.com/t67y110v/driver/internal/driver/parser"
	// "github.com/t67y110v/driver/internal/driver/scale"
	pb "github.com/t67y110v/driver/pkg/driver"
	"google.golang.org/grpc"
)

func (s *ScalesHandler) Example() {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", s.config.ScaleIP, s.config.ScalePort))
	if err != nil {
		return
	}
	defer conn.Close()
	client := pb.NewApiCallerScaleClient(conn)

	s.SetZero(client)
	s.SetTareValue(client)
	s.SetTare(client)
	s.ScalesMessageOutChannel(client)
	s.GetState(client)
	s.GetInstantWeight(client)

}

// func (s *ScalesHandler) ScalesMessageOutChannel(stream driver.ApiCallerScale_ScalesMessageOutChannelServer) error {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	for {
// 		{
// 			stream.Send(&driver.ResponseScale{
// 				Message: string(s.scale.MakeMessage([]byte{scale.Commands["CMD_SET_ZERO"]})),
// 			})

// 			req, err := stream.Recv()
// 			if err == io.EOF {
// 				return nil
// 			}
// 			if err != nil {
// 				return err
// 			}
// 			msg := []byte(req.Message)

// 			resp := &model.SetZeroResponse{}
// 			err = resp.Read(msg)
// 			if err != nil {
// 				continue
// 			}
// 			parser.FillResponseStruct(resp.Raw(), reflect.ValueOf(&resp).Elem())
// 			if resp.Response.Raw().Meta.CommandCode == scale.Commands["CMD_ACK_SET"] {
// 				s.logger.Infoln("weight value set to zero")
// 			}
// 		}

// 		{
// 			stream.Send(&driver.ResponseScale{
// 				Message: string(s.scale.MakeMessagefForSetValue([]byte{scale.Commands["CMD_SET_TARE"]}, []byte{12})),
// 			})

// 			req, err := stream.Recv()
// 			if err == io.EOF {
// 				return nil
// 			}
// 			if err != nil {
// 				return err
// 			}
// 			msg := []byte(req.Message)
// 			resp := &model.SetTareResponse{}
// 			err = resp.Read(msg)
// 			if err != nil {
// 				continue
// 			}
// 			parser.FillResponseStruct(resp.Raw(), reflect.ValueOf(&resp).Elem())
// 			s.logger.Info("weight value set to input value")
// 		}
// 		{

// 		}

// 	}

// }
func (h *ScalesHandler) ScalesMessageOutChannel(client pb.ApiCallerScaleClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.ScalesMessageOutChannel(ctx)
	if err != nil {
		h.logger.Fatalf("client.ScalesMessageOutChannel failed: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				h.logger.Fatalf("client.RouteChat failed: %v", err)
			}
			h.logger.Printf("Got message %s  with error : %s type : %s subtype : %s", in.Message, in.Error, in.Type, in.Subtype)
		}
	}()
	// for _, note := range notes {
	// 	if err := stream.Send(note); err != nil {
	// 		log.Fatalf("client.RouteChat: stream.Send(%v) failed: %v", note, err)
	// 	}
	// }
	stream.CloseSend()
	<-waitc
}
func (h *ScalesHandler) SetTare(client pb.ApiCallerScaleClient) {
	h.logger.Printf("Setting tare for scale (%s)", h.config.ScaleIP)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.SetTare(ctx, &pb.Empty{})
	if err != nil {
		h.logger.Fatalf("client.GetFeature failed: %v", err)
	}
	h.logger.Println(resp)

}
func (h *ScalesHandler) SetTareValue(client pb.ApiCallerScaleClient) {
	h.logger.Printf("Setting tare value for scale (%s)", h.config.ScaleIP)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.SetTareValue(ctx, &pb.RequestTareValue{})
	if err != nil {
		h.logger.Fatalf("client.GetFeature failed: %v", err)
	}
	h.logger.Println(resp)

}
func (h *ScalesHandler) SetZero(client pb.ApiCallerScaleClient) {
	h.logger.Printf("Setting zero value for scale (%s)", h.config.ScaleIP)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.SetZero(ctx, &pb.Empty{})
	if err != nil {
		h.logger.Fatalf("client.GetFeature failed: %v", err)
	}
	h.logger.Println(resp)
}
func (h *ScalesHandler) GetInstantWeight(client pb.ApiCallerScaleClient) {
	h.logger.Printf("Getting instant weight for scale (%s)", h.config.ScaleIP)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	weight, err := client.GetInstantWeight(ctx, &pb.Empty{})
	if err != nil {
		h.logger.Fatalf("client.GetFeature failed: %v", err)
	}
	h.logger.Println(weight)

}

func (h *ScalesHandler) GetState(client pb.ApiCallerScaleClient) {
	h.logger.Printf("Getting State for scale (%s)", h.config.ScaleIP)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	state, err := client.GetState(ctx, &pb.Empty{})
	if err != nil {
		h.logger.Fatalf("client.GetFeature failed: %v", err)
	}
	h.logger.Println(state)

}
