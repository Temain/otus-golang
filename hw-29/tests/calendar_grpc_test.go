package main

import (
	"context"
	"fmt"
	event "github.com/Temain/otus-golang/hw-29/pkg/proto"
	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

type calendarGrpcTest struct {
	ctx        context.Context
	clientConn *grpc.ClientConn
	client     event.EventServiceClient
	list       []event.EventMessage
	found      *event.EventMessage
}

func (test *calendarGrpcTest) connect(*messages.Pickle) {
	test.ctx, _ = context.WithTimeout(context.Background(), 5*time.Minute)
	cc, err := grpc.Dial("grpc_api:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	test.client = event.NewEventServiceClient(cc)
	test.clientConn = cc
}

func (test *calendarGrpcTest) close(*messages.Pickle, error) {
	err := test.clientConn.Close()
	if err != nil {
		fmt.Errorf("error on close connection: %v", err)
	}
}

func (test *calendarGrpcTest) iCallListMethod() error {
	stream, err := test.client.List(test.ctx, &event.ListRequest{})
	if err != nil {
		err = fmt.Errorf("error on list events: %v", err)
	}

	if stream == nil {
		err = fmt.Errorf("wrong result on list of events: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			fmt.Errorf("recv EOF")
			break
		}
		if err != nil {
			err = fmt.Errorf("error on receive list of events: %v", err)
		}
		if msg == nil {
			err = fmt.Errorf("received message is empty")
		}
		fmt.Errorf("received message: %v", &msg)
		test.list = append(test.list, *msg)
	}

	return nil
}

func (test *calendarGrpcTest) theListResultShouldBeNonEmpty() error {
	if len(test.list) == 0 {
		return fmt.Errorf("result of list method is empty")
	}
	return nil
}

func (test *calendarGrpcTest) iCallSearchMethod() error {
	sample := time.Date(2020, 04, 25, 22, 00, 00, 00, time.UTC)
	created, err := ptypes.TimestampProto(sample)
	if err != nil {
		log.Fatalf("wrong event date: %v", err)
	}
	response, err := test.client.Search(test.ctx, &event.SearchRequest{Date: created})
	if err != nil {
		log.Fatalf("error on search event: %v", err)
	}

	test.found = response.Event

	return nil
}

func (test *calendarGrpcTest) theSearchResultShouldBeNonEmpty() error {
	if test.found == nil {
		return fmt.Errorf("result of search method is empty")
	}
	return nil
}

func FeatureContextGrpc(s *godog.Suite) {
	testGrpc := new(calendarGrpcTest)

	s.BeforeScenario(testGrpc.connect)

	s.Step(`^I call list method$`, testGrpc.iCallListMethod)
	s.Step(`^The result should be non empty$`, testGrpc.theListResultShouldBeNonEmpty)

	s.Step(`^I call search method$`, testGrpc.iCallSearchMethod)
	s.Step(`^Method should return 1 event$`, testGrpc.theSearchResultShouldBeNonEmpty)

	s.AfterScenario(testGrpc.close)
}
