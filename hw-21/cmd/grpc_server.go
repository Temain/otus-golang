package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	c "github.com/Temain/otus-golang/hw-21/internal/calendar"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	e "github.com/Temain/otus-golang/hw-21/internal/calendar/entities"
	i "github.com/Temain/otus-golang/hw-21/internal/calendar/interfaces"
	p "github.com/Temain/otus-golang/hw-21/internal/proto"

	"github.com/spf13/cobra"
)

var index int64 = 1

type EventServer struct {
	Calendar i.Calendar
}

func (s EventServer) List(request *p.ListRequest, stream p.EventService_ListServer) error {
	events, err := s.Calendar.List()
	if err != nil {
		return status.Error(codes.Internal, "error on list events")
	}

	for _, event := range events {
		msg, err := mapEventToMessage(event)
		if err != nil {
			return status.Error(codes.Internal, "error on map event to response message")
		}
		if err := stream.Send(msg); err != nil {
			return err
		}
	}

	return nil
}

func (s EventServer) Search(context context.Context, request *p.SearchRequest) (*p.SearchResponse, error) {
	response := &p.SearchResponse{}
	created, err := ptypes.Timestamp(request.Date)
	if err != nil {
		return response, status.Error(codes.InvalidArgument, "wrong search date")
	}
	event, _ := s.Calendar.Search(created)
	if event == nil {
		return response, nil
	}

	message, err := mapEventToMessage(*event)
	if err != nil {
		return response, status.Error(codes.Internal, "error on map event to response message")
	}
	response.Event = message
	return response, nil
}

func (s EventServer) Add(context context.Context, request *p.AddRequest) (*p.AddResponse, error) {
	defer func() { index++ }()
	response := &p.AddResponse{}
	event, err := mapMessageToEvent(*request.Event)
	if err != nil {
		return response, status.Error(codes.Internal, "error on map event to response message")
	}

	event.Id = index
	err = s.Calendar.Add(event)
	if err != nil {
		return response, status.Error(codes.Internal, fmt.Sprintf("error on add new event: %v", err))
	}

	response.Success = true
	return response, nil
}

func (s EventServer) Update(context context.Context, request *p.UpdateRequest) (*p.UpdateResponse, error) {
	response := &p.UpdateResponse{}
	event, err := mapMessageToEvent(*request.Event)
	if err != nil {
		return response, status.Error(codes.Internal, "error on map message to event")
	}

	err = s.Calendar.Update(event)
	if err != nil {
		return response, status.Error(codes.Internal, fmt.Sprintf("error on update event: %v", err))
	}

	response.Success = true
	return response, nil
}

func (s EventServer) Delete(context context.Context, request *p.DeleteRequest) (*p.DeleteResponse, error) {
	response := &p.DeleteResponse{}
	err := s.Calendar.Delete(request.Id)
	if err != nil {
		return response, status.Error(codes.Internal, fmt.Sprintf("error on delete event: %v", err))
	}

	response.Success = true
	return response, nil
}

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

		// calendar with some events
		calendar := c.NewCalendar()
		event1 := &e.Event{
			Id:          index,
			Title:       "Morning coffee",
			Description: "The most important event of the day",
			Created:     time.Date(2020, 04, 22, 10, 00, 00, 00, time.UTC),
		}
		_ = calendar.Add(event1)
		index++

		event2 := &e.Event{
			Id:          index,
			Title:       "Evening tea",
			Description: "Not bad",
			Created:     time.Date(2020, 04, 22, 22, 00, 00, 00, time.UTC),
		}
		_ = calendar.Add(event2)
		index++

		eventServer := &EventServer{Calendar: calendar}

		p.RegisterEventServiceServer(grpcServer, eventServer)
		_ = grpcServer.Serve(listen)
	},
}

func mapEventToMessage(event e.Event) (*p.EventMessage, error) {
	timestamp, err := ptypes.TimestampProto(event.Created)
	if err != nil {
		return nil, errors.New("wrong event time")
	}
	msg := &p.EventMessage{
		Id:          event.Id,
		Title:       event.Title,
		Description: event.Description,
		Created:     timestamp,
	}
	return msg, nil
}

func mapMessageToEvent(msg p.EventMessage) (*e.Event, error) {
	created, err := ptypes.Timestamp(msg.Created)
	if err != nil {
		return nil, errors.New("wrong event time in message")
	}

	event := &e.Event{
		Id:          msg.Id,
		Title:       msg.Title,
		Description: msg.Description,
		Created:     created,
	}
	return event, nil
}
