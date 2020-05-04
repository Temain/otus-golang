package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"google.golang.org/grpc"

	"github.com/Temain/otus-golang/project/pkg/proto"
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
	client     proto.RotationServiceClient
	//sampleEvent        *event.EventMessage
	//listResult         []event.EventMessage
	//searchResult       *event.EventMessage
	//addResult          bool
	//addDuplicateResult bool
	//updateResult       bool
	//updateNotExists    bool
	//deleteResult       bool
	//deleteNotExists    bool
}

func (test *rotationGrpcTest) connect(*messages.Pickle) {
	test.ctx, _ = context.WithTimeout(context.Background(), 5*time.Minute)
	cc, err := grpc.Dial(grpcListen, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	test.client = proto.NewRotationServiceClient(cc)
	test.clientConn = cc
}

func (test *rotationGrpcTest) close(*messages.Pickle, error) {
	err := test.clientConn.Close()
	if err != nil {
		fmt.Errorf("error on close connection: %v", err)
	}
}

func (test *rotationGrpcTest) iCallAddMethod() error {
	return godog.ErrPending
}

func (test *rotationGrpcTest) methodShouldReturnSuccessResult() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	testGrpc := new(rotationGrpcTest)

	s.BeforeScenario(testGrpc.connect)

	s.Step(`^I call add method$`, testGrpc.iCallAddMethod)
	s.Step(`^Method should return success result$`, testGrpc.methodShouldReturnSuccessResult)

	s.AfterScenario(testGrpc.close)
}
