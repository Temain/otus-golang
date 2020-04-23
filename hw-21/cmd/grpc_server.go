package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	c "github.com/Temain/otus-golang/hw-21/internal/calendar"
	p "github.com/Temain/otus-golang/hw-21/internal/proto"

	"github.com/spf13/cobra"
)

type EventServer struct {
}

func (s EventServer) List(req *p.ListRequest, stream p.EventService_ListServer) error {
	calendar := c.NewCalendar()

	// Test
	event := &c.Event{
		Id:          1,
		Title:       "Morning coffee",
		Description: "The most important event of the day",
		Created:     time.Now(),
	}
	_ = calendar.Add(event)

	events, err := calendar.List()
	if err != nil {
		log.Fatalf("error on list events: %v", err)
	}

	for _, event := range events {
		timestamp, err := ptypes.TimestampProto(event.Created)
		if err != nil {
			log.Fatalf("wrong event time: %v", err)
		}
		msg := &p.EventMessage{
			Id:          int64(event.Id),
			Title:       event.Title,
			Description: event.Description,
			Created:     timestamp,
		}

		if err := stream.Send(msg); err != nil {
			return err
		}
	}
	return nil
}

func (s EventServer) Search(context.Context, *p.SearchRequest) (*p.SearchResponse, error) {
	panic("implement me")
}

func (s EventServer) Add(context.Context, *p.AddRequest) (*p.AddResponse, error) {
	panic("implement me")
}

func (s EventServer) Update(context.Context, *p.UpdateRequest) (*p.UpdateResponse, error) {
	panic("implement me")
}

func (s EventServer) Delete(context.Context, *p.DeleteRequest) (*p.DeleteResponse, error) {
	panic("implement me")
}

//var i int64
//
//func (s EventServer) SendMessage(ctx context.Context, msg *event.EventMessage) (*event.EventMessage, error) {
//	defer func() { i++ }()
//	if msg.Title == "" {
//		return nil, status.Error(codes.InvalidArgument, "no empty string")
//	}
//	return &event.EventMessage{
//		Id:          i,
//		Title:       msg.Title,
//		Description: msg.Description,
//		Created:     ptypes.TimestampNow(),
//	}, nil
//}

var GrpcServerCmd = &cobra.Command{
	Use:   "grpc_server",
	Short: "run grpc server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running gRPC server...")

		listen, err := net.Listen("tcp", "0.0.0.0:50051")
		if err != nil {
			log.Fatalf("failed to listen %v", err)
		}

		grpcServer := grpc.NewServer()
		reflection.Register(grpcServer)
		p.RegisterEventServiceServer(grpcServer, EventServer{})
		_ = grpcServer.Serve(listen)
	},
}
