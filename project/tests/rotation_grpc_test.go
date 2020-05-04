package main

import (
	"context"
	"os"

	pr "github.com/Temain/otus-golang/project/pkg/proto"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"google.golang.org/grpc"
)

var grpcListen = os.Getenv("TESTS_GRPC_API")

func init() {
	if grpcListen == "" {
		grpcListen = "grpc_api:50052"
	}
}

type rotationGrpcTest struct {
	ctx        context.Context
	clientConn *grpc.ClientConn
	client     pr.RotationServiceClient
	// sampleEvent        *event.EventMessage
	// listResult         []event.EventMessage
	// searchResult       *event.EventMessage
	// addResult          bool
	addDuplicateResult bool
	updateResult       bool
	updateNotExists    bool
	deleteResult       bool
	deleteNotExists    bool
}

func (test *rotationGrpcTest) connect(*messages.Pickle) {
	// test.ctx, _ = context.WithTimeout(context.Background(), 5*time.Minute)
	//cc, err := grpc.Dial(grpcListen, grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("could not connect: %v", err)
	//}
	//test.client = proto.NewEventServiceClient(cc)
	//test.clientConn = cc
	//
	//sampleTime := time.Date(2020, 04, 22, 10, 00, 00, 00, time.UTC)
	//created, err := ptypes.TimestampProto(sampleTime)
	//if err != nil {
	//	log.Fatalf("wrong event date: %v", err)
	//}
	//test.sampleEvent = &proto.EventMessage{
	//	Id:          1,
	//	Title:       "Sample event title",
	//	Description: "Sample event description",
	//	Created:     created,
	//}
}

func (test *rotationGrpcTest) close(*messages.Pickle, error) {
	//err := test.clientConn.Close()
	//if err != nil {
	//	fmt.Errorf("error on close connection: %v", err)
	//}
}

func (test *rotationGrpcTest) iCallAddMethod() error {
	//request := &proto.AddBannerRequest{}
	//response, err := test.client.Add(test.ctx, request)
	//if err != nil {
	//	return fmt.Errorf("error on add event: %v", err)
	//}

	// test.addResult = response.Success

	return nil
}

func (test *rotationGrpcTest) theAddResultShouldBeSuccess() error {
	//if !test.addResult {
	//	return fmt.Errorf("new event not added")
	//}
	return nil
}

func FeatureContext(s *godog.Suite) {
	testGrpc := new(rotationGrpcTest)

	s.BeforeScenario(testGrpc.connect)

	s.Step(`^I call add method$`, testGrpc.iCallAddMethod)
	s.Step(`^Method should return success result$`, testGrpc.theAddResultShouldBeSuccess)

	s.AfterScenario(testGrpc.close)
}
