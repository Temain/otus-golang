package cmd

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	event "github.com/Temain/otus-golang/hw-21/internal/proto"

	"github.com/spf13/cobra"
)

type EventServer struct {
}

var i int64

func (s EventServer) SendMessage(ctx context.Context, msg *event.EventMessage) (*event.EventMessage, error) {
	defer func() { i++ }()
	if msg.Title == "" {
		return nil, status.Error(codes.InvalidArgument, "no empty string")
	}
	return &event.EventMessage{
		Id:          i,
		Title:       msg.Title,
		Description: msg.Description,
		Date:        ptypes.TimestampNow(),
	}, nil
}

var GrpcServerCmd = &cobra.Command{
	Use:   "grpc_server",
	Short: "run grpc server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running GRPC server...")

		lis, err := net.Listen("tcp", "0.0.0.0:50051")
		if err != nil {
			log.Fatalf("failed to listen %v", err)
		}

		grpcServer := grpc.NewServer()
		reflection.Register(grpcServer)
		event.RegisterEventServiceServer(grpcServer, EventServer{})
		_ = grpcServer.Serve(lis)
	},
}
