package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/jmoiron/sqlx"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Temain/otus-golang/hw-22/internal/calendar"
	"github.com/Temain/otus-golang/hw-22/internal/calendar/entities"
	interfaces "github.com/Temain/otus-golang/hw-22/internal/calendar/interfaces"
	"github.com/Temain/otus-golang/hw-22/internal/configer"
	"github.com/Temain/otus-golang/hw-22/internal/logger"
	"github.com/golang/protobuf/ptypes"

	"github.com/Temain/otus-golang/hw-22/internal/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var logr *logrus.Logger
var index int64 = 1

type EventServer struct {
	Calendar interfaces.EventAdapter
}

func StartGrpcServer() error {
	cfg := configer.ReadConfig()
	logr = logger.NewLogger(cfg.LogFile, cfg.LogLevel)
	addr := fmt.Sprintf("0.0.0.0%s", cfg.GrpcListen)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	db, err := sqlx.Open("pgx", cfg.PostgresDsn)
	if err != nil {
		log.Fatalf("connection to database failed: %v", err)
	}

	calendar, err := calendar.NewPostgresCalendar(db)
	if err != nil {
		log.Fatalf("unable to create calendar: %v", err)
	}
	log.Println("connected to database")

	eventServer := &EventServer{Calendar: calendar}
	proto.RegisterEventServiceServer(grpcServer, eventServer)
	return grpcServer.Serve(listen)
}

func (s EventServer) List(request *proto.ListRequest, stream proto.EventService_ListServer) error {
	logr.Info("received list request")
	events, err := s.Calendar.List(context.Background())
	if err != nil {
		log.Printf("error on list events: %v", err)
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

func (s EventServer) Search(ctx context.Context, request *proto.SearchRequest) (*proto.SearchResponse, error) {
	logr.Info("received search request")
	response := &proto.SearchResponse{}
	created, err := ptypes.Timestamp(request.Date)
	if err != nil {
		return response, status.Error(codes.InvalidArgument, "wrong search date")
	}
	event, _ := s.Calendar.Search(ctx, created)
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

func (s EventServer) Add(ctx context.Context, request *proto.AddRequest) (*proto.AddResponse, error) {
	logr.Info("received add request")
	defer func() { index++ }()
	response := &proto.AddResponse{}
	event, err := mapMessageToEvent(*request.Event)
	if err != nil {
		return response, status.Error(codes.Internal, "error on map event to response message")
	}

	event.Id = index
	err = s.Calendar.Add(ctx, event)
	if err != nil {
		return response, status.Error(codes.Internal, fmt.Sprintf("error on add new event: %v", err))
	}

	response.Success = true
	return response, nil
}

func (s EventServer) Update(ctx context.Context, request *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	logr.Info("received update request")
	response := &proto.UpdateResponse{}
	event, err := mapMessageToEvent(*request.Event)
	if err != nil {
		return response, status.Error(codes.Internal, "error on map message to event")
	}

	err = s.Calendar.Update(ctx, event)
	if err != nil {
		return response, status.Error(codes.Internal, fmt.Sprintf("error on update event: %v", err))
	}

	response.Success = true
	return response, nil
}

func (s EventServer) Delete(ctx context.Context, request *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	logr.Info("received delete request")
	response := &proto.DeleteResponse{}
	err := s.Calendar.Delete(ctx, request.Id)
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
