package api

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Temain/otus-golang/hw-21/internal/calendar"
	"github.com/Temain/otus-golang/hw-21/internal/calendar/entities"
	interfaces "github.com/Temain/otus-golang/hw-21/internal/calendar/interfaces"
	"github.com/Temain/otus-golang/hw-21/internal/configer"
	"github.com/Temain/otus-golang/hw-21/internal/logger"
	"github.com/Temain/otus-golang/hw-21/internal/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var logr *logrus.Logger

type EventServer struct {
	Calendar interfaces.EventAdapter
}

func StartGrpcServer() error {
	cfg := configer.ReadConfig()
	logr = logger.NewLogger(cfg.LogFile, cfg.LogLevel)
	addr := fmt.Sprintf("0.0.0.0%s", cfg.GrpcListen)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	// calendar with some events
	calendar := calendar.NewMemoryCalendar()
	event1 := &entities.Event{
		Title:       "Morning coffee",
		Description: "The most important event of the day",
		Created:     time.Date(2020, 04, 22, 10, 00, 00, 00, time.UTC),
	}
	_ = calendar.Add(event1)

	event2 := &entities.Event{
		Title:       "Evening tea",
		Description: "Not bad",
		Created:     time.Date(2020, 04, 22, 22, 00, 00, 00, time.UTC),
	}
	_ = calendar.Add(event2)

	eventServer := &EventServer{Calendar: calendar}

	proto.RegisterEventServiceServer(grpcServer, eventServer)
	return grpcServer.Serve(listen)
}

func (s EventServer) List(request *proto.ListRequest, stream proto.EventService_ListServer) error {
	logr.Info("received list request")
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

func (s EventServer) Search(context context.Context, request *proto.SearchRequest) (*proto.SearchResponse, error) {
	logr.Info("received search request")
	response := &proto.SearchResponse{}
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

func (s EventServer) Add(context context.Context, request *proto.AddRequest) (*proto.AddResponse, error) {
	logr.Info("received add request")
	response := &proto.AddResponse{}
	event, err := mapMessageToEvent(*request.Event)
	if err != nil {
		return response, status.Error(codes.Internal, "error on map event to response message")
	}

	err = s.Calendar.Add(event)
	if err != nil {
		return response, status.Error(codes.Internal, fmt.Sprintf("error on add new event: %v", err))
	}

	response.Success = true
	return response, nil
}

func (s EventServer) Update(context context.Context, request *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	logr.Info("received update request")
	response := &proto.UpdateResponse{}
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

func (s EventServer) Delete(context context.Context, request *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	logr.Info("received delete request")
	response := &proto.DeleteResponse{}
	err := s.Calendar.Delete(request.Id)
	if err != nil {
		return response, status.Error(codes.Internal, fmt.Sprintf("error on delete event: %v", err))
	}

	response.Success = true
	return response, nil
}

func mapEventToMessage(event entities.Event) (*proto.EventMessage, error) {
	timestamp, err := ptypes.TimestampProto(event.Created)
	if err != nil {
		return nil, errors.New("wrong event time")
	}
	msg := &proto.EventMessage{
		Id:          event.Id,
		Title:       event.Title,
		Description: event.Description,
		Created:     timestamp,
	}
	return msg, nil
}

func mapMessageToEvent(msg proto.EventMessage) (*entities.Event, error) {
	created, err := ptypes.Timestamp(msg.Created)
	if err != nil {
		return nil, errors.New("wrong event time in message")
	}

	event := &entities.Event{
		Id:          msg.Id,
		Title:       msg.Title,
		Description: msg.Description,
		Created:     created,
	}
	return event, nil
}
