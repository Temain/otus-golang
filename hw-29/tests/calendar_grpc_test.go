package main

import (
	"context"
	event "github.com/Temain/otus-golang/hw-29/pkg/proto"
	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"log"
	"time"
)

type calendarGrpcTest struct {
	ctx        context.Context
	// clientConn *grpc.ClientConn
	client     event.EventServiceClient
	list       []event.EventMessage
}

func (test *calendarGrpcTest) connectGrpc(*messages.Pickle) {
	test.ctx, _ = context.WithTimeout(context.Background(), 5*time.Minute)
	cc, err := grpc.Dial("grpc_api:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	test.client = event.NewEventServiceClient(cc)
}

func (test *calendarGrpcTest) iCallListMethod() error {
	/*stream, err := test.client.List(test.ctx, &event.ListRequest{})
	if err != nil {
		err = fmt.Errorf("error on list events: %v", err)
	}

	if stream == nil {
		log.Fatalf("!!!!!! NULLLLLLLLLLLLLLLLLLLLLLL")
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			err = fmt.Errorf("error on receive list of events: %v", err)
		}
		if msg == nil {
			err = fmt.Errorf("received message is empty")
		}
		test.list = append(test.list, *msg)
	}*/

	sample := time.Date(2020, 04, 25, 22, 00, 00, 00, time.UTC)
	created, err := ptypes.TimestampProto(sample)
	if err != nil {
		log.Fatalf("wrong event date: %v", err)
	}
	response, err := test.client.Search(test.ctx, &event.SearchRequest{Date: created})
	if err != nil {
		log.Fatalf("error on search event: %v", err)
	}

	msg := response.Event
	if msg == nil {
		log.Println("event not found")
	}

	log.Printf("found event: %v\n", msg)

	return nil
}

func (test *calendarGrpcTest) theResultShouldBeNonEmpty() error {
	return godog.ErrPending
}

func FeatureContextGrpc(s *godog.Suite) {
	testGrpc := new(calendarGrpcTest)

	s.BeforeScenario(testGrpc.connectGrpc)

	s.Step(`^I call list method$`, testGrpc.iCallListMethod)
	s.Step(`^The result should be non empty$`, testGrpc.theResultShouldBeNonEmpty)
}
