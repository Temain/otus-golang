package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	event "github.com/Temain/otus-golang/hw-29/pkg/proto"
	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

type calendarGrpcTest struct {
	ctx          context.Context
	clientConn   *grpc.ClientConn
	client       event.EventServiceClient
	sampleEvent  *event.EventMessage
	listResult   []event.EventMessage
	searchResult *event.EventMessage
	addResult    bool
	updateResult bool
	deleteResult bool
}

func (test *calendarGrpcTest) connect(*messages.Pickle) {
	test.ctx, _ = context.WithTimeout(context.Background(), 5*time.Minute)
	cc, err := grpc.Dial("grpc_api:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	test.client = event.NewEventServiceClient(cc)
	test.clientConn = cc

	sampleTime := time.Date(2020, 04, 22, 10, 00, 00, 00, time.UTC)
	created, err := ptypes.TimestampProto(sampleTime)
	if err != nil {
		log.Fatalf("wrong event date: %v", err)
	}
	test.sampleEvent = &event.EventMessage{
		Id:          1,
		Title:       "Sample event title",
		Description: "Sample event description",
		Created:     created,
	}
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
		test.listResult = append(test.listResult, *msg)
	}

	return nil
}

func (test *calendarGrpcTest) theListResultShouldBeNonEmpty() error {
	if len(test.listResult) == 0 {
		return fmt.Errorf("result of list method is empty")
	}
	return nil
}

func (test *calendarGrpcTest) iCallSearchMethod() error {
	created := test.sampleEvent.Created
	response, err := test.client.Search(test.ctx, &event.SearchRequest{Date: created})
	if err != nil {
		log.Fatalf("error on search event: %v", err)
	}

	test.searchResult = response.Event

	return nil
}

func (test *calendarGrpcTest) theSearchResultShouldBeNonEmpty() error {
	if test.searchResult == nil {
		return fmt.Errorf("result of search method is empty")
	}
	return nil
}

func (test *calendarGrpcTest) iCallAddMethod() error {
	request := &event.AddRequest{Event: test.sampleEvent}
	response, err := test.client.Add(test.ctx, request)
	if err != nil {
		log.Fatalf("error on add event: %v", err)
	}

	test.addResult = response.Success

	return nil
}

func (test *calendarGrpcTest) theAddResultShouldBeSuccess() error {
	if !test.addResult {
		return fmt.Errorf("new event not added")
	}
	return nil
}

func (test *calendarGrpcTest) iCallUpdateMethod() error {
	request := &event.UpdateRequest{Event: test.sampleEvent}
	response, err := test.client.Update(test.ctx, request)
	if err != nil {
		log.Fatalf("error on update event: %v", err)
	}

	test.updateResult = response.Success

	return nil
}

func (test *calendarGrpcTest) theUpdateResultShouldBeSuccess() error {
	if !test.updateResult {
		return fmt.Errorf("event not updated")
	}
	return nil
}

func (test *calendarGrpcTest) iCallDeleteMethod() error {
	request := &event.DeleteRequest{Id: 1}
	response, err := test.client.Delete(test.ctx, request)
	if err != nil {
		log.Fatalf("error on delete event: %v", err)
	}

	test.deleteResult = response.Success

	return nil
}

func (test *calendarGrpcTest) theDeleteResultShouldBeSuccess() error {
	if !test.deleteResult {
		return fmt.Errorf("event not deleted")
	}
	return nil
}

func FeatureContextGrpc(s *godog.Suite) {
	testGrpc := new(calendarGrpcTest)

	s.BeforeScenario(testGrpc.connect)

	s.Step(`^I call add method$`, testGrpc.iCallAddMethod)
	s.Step(`^Method should return success result$`, testGrpc.theAddResultShouldBeSuccess)

	s.Step(`^I call list method$`, testGrpc.iCallListMethod)
	s.Step(`^The result should be non empty$`, testGrpc.theListResultShouldBeNonEmpty)

	s.Step(`^I call search method$`, testGrpc.iCallSearchMethod)
	s.Step(`^Method should return 1 event$`, testGrpc.theSearchResultShouldBeNonEmpty)

	s.Step(`^I call update method$`, testGrpc.iCallUpdateMethod)
	s.Step(`^Method should return success result$`, testGrpc.theUpdateResultShouldBeSuccess)

	s.Step(`^I call delete method$`, testGrpc.iCallDeleteMethod)
	s.Step(`^Method should return success result$`, testGrpc.theDeleteResultShouldBeSuccess)

	s.AfterScenario(testGrpc.close)
}
