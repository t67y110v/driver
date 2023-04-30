package handlers

import (
	"context"
	"fmt"
	"io"
	"math/big"
	"reflect"
	"time"

	"github.com/t67y110v/driver/internal/driver/converter"
	"github.com/t67y110v/driver/internal/driver/model"
	"github.com/t67y110v/driver/internal/driver/parser"
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
	s.SetTareValue(client, &pb.RequestTareValue{})
	s.SetTare(client)
	s.ScalesMessageOutChannel(client)
	s.GetState(client)
	s.GetInstantWeight(client)

}

func (h *ScalesHandler) ScalesMessageOutChannel(client pb.ApiCallerScaleClient) {
	h.mu.Lock()
	defer h.mu.Unlock()
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
			msg := []byte(in.Message)
			str := &model.Response{}
			err = str.Read(msg)
			if err != nil {
				return
			}
			parser.MakeResponse(str.Raw(), reflect.ValueOf(&str).Elem())

			if str.Raw().Meta.CommandCode != converter.Responses["CMD_ERROR"] {
				h.logger.Printf("Code :%s\n", str.Raw().Meta.CommandCode)
			} else {
				h.logger.Println("command execution error: unable to install tare")
			}
		}
	}()
	stream.CloseSend()
	<-waitc
}
func (h *ScalesHandler) SetTare(client pb.ApiCallerScaleClient) {
	h.logger.Printf("Setting tare for scale (%s)", h.config.ScaleIP)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.SetTare(ctx, &pb.Empty{})
	if err != nil {
		h.logger.Fatalf("client.SetTare failed: %v", err)
	}

	msg := []byte(resp.Error)
	str := &model.SetTareResponse{}
	err = str.Read(msg)
	if err != nil {
		return
	}
	parser.MakeResponse(str.Raw(), reflect.ValueOf(&str).Elem())

	if str.Response.Raw().Meta.CommandCode == converter.Responses["CMD_ACK_SET_TARE"] {
		h.logger.Println("Set tare command completed successfully")
	} else if str.Response.Raw().Meta.CommandCode == converter.Responses["CMD_NACK_TARE"] {
		h.logger.Println("command execution error: unable to install tare")
	}

}
func (h *ScalesHandler) SetTareValue(client pb.ApiCallerScaleClient, value *pb.RequestTareValue) {
	h.logger.Printf("Setting tare value for scale (%s)", h.config.ScaleIP)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var tareValue int64 = 1500
	var tareValueBytes []byte = big.NewInt(tareValue).Bytes()
	req := &pb.RequestTareValue{Message: string(converter.MakeMessagefForSetValue([]byte{converter.Commands["CMD_SET_TARE"]}, tareValueBytes))}

	resp, err := client.SetTareValue(ctx, req)
	if err != nil {
		h.logger.Fatalf("client.SetTareValue failed: %v", err)
	}
	msg := []byte(resp.Error)
	str := &model.SetTareResponse{}
	err = str.Read(msg)
	if err != nil {
		return
	}
	parser.MakeResponse(str.Raw(), reflect.ValueOf(&str).Elem())
	if str.Response.Raw().Meta.CommandCode == converter.Responses["CMD_ACK_SET_TARE"] {
		h.logger.Printf("Set tare command completed successfully, value that is set to %d\n", tareValue)
	} else if str.Response.Raw().Meta.CommandCode == converter.Responses["CMD_NACK_TARE"] {
		h.logger.Println("command execution error: unable to install tare")
	}

}
func (h *ScalesHandler) SetZero(client pb.ApiCallerScaleClient) {
	h.logger.Printf("Setting zero value for scale (%s)", h.config.ScaleIP)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.SetZero(ctx, &pb.Empty{})
	if err != nil {
		h.logger.Fatalf("client.SetZero failed: %v", err)
	}
	msg := []byte(resp.Error)
	str := &model.SetTareZeroResponse{}
	err = str.Read(msg)
	if err != nil {
		return
	}
	parser.MakeResponse(str.Raw(), reflect.ValueOf(&str).Elem())
	if str.Response.Raw().Meta.CommandCode == converter.Responses["CMD_ACK_SET"] {
		h.logger.Println("Set tare command completed successfully, value that is set to zero")
	} else if str.Response.Raw().Meta.CommandCode == converter.Responses["CMD_ERROR"] {
		h.logger.Println("command execution error: unable to install tare")
	}

}
func (h *ScalesHandler) GetInstantWeight(client pb.ApiCallerScaleClient) {
	h.logger.Printf("Getting instant weight for scale (%s)", h.config.ScaleIP)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.GetInstantWeight(ctx, &pb.Empty{})
	if err != nil {
		h.logger.Fatalf("client.GetInstantWeight failed: %v", err)
	}
	msg := []byte(resp.Message)
	str := &model.GetMassaResponse{}
	err = str.Read(msg)
	if err != nil {
		return
	}
	parser.MakeResponse(str.Raw(), reflect.ValueOf(&str).Elem())
	if str.Response.Raw().Meta.CommandCode == converter.Responses["CMD_ACK_MASSA"] {
		h.logger.Printf("Get massa command completed successfully, weight = %d, stable = %d, net = %d, zero =%d", str.Weight, str.Stable, str.Net, str.Zero)
	} else if str.Response.Raw().Meta.CommandCode == converter.Responses["CMD_ERROR"] {
		h.logger.Println("command execution error: unable to install tare")
	}

}

func (h *ScalesHandler) GetState(client pb.ApiCallerScaleClient) {
	h.logger.Printf("Getting State for scale (%s)", h.config.ScaleIP)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	state, err := client.GetState(ctx, &pb.Empty{})
	if err != nil {
		h.logger.Fatalf("client.GetState failed: %v", err)
	}

	msg := []byte(state.Message)
	str := &model.Response{}
	err = str.Read(msg)
	if err != nil {
		return
	}
	parser.MakeResponse(str.Raw(), reflect.ValueOf(&str).Elem())
	if str.Raw().Meta.CommandCode == converter.Responses["CMD_ACK_STATE"] {
		h.logger.Println("Get massa command completed successfully")
	} else if str.Raw().Meta.CommandCode == converter.Responses["CMD_ERROR"] {
		h.logger.Println("command execution error: unable to install tare")
	}

}
