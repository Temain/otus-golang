package tests

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	proto "github.com/Temain/otus-golang/project/pkg/proto"
	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"google.golang.org/grpc"
)

var grpcListen = os.Getenv("TESTS_GRPC_API")

func init() {
	if grpcListen == "" {
		grpcListen = "grpc_api:50051"
	}
}

type rotationGrpcTest struct {
	ctx        context.Context
	clientConn *grpc.ClientConn
	client     proto.RotationServiceClient
	//sampleEvent        *event.EventMessage
	//listResult         []event.EventMessage
	//searchResult       *event.EventMessage
	addResult          bool
	addDuplicateResult bool
	updateResult       bool
	updateNotExists    bool
	deleteResult       bool
	deleteNotExists    bool
}

func (test *rotationGrpcTest) connect(*messages.Pickle) {
	test.ctx, _ = context.WithTimeout(context.Background(), 5*time.Minute)
	cc, err := grpc.Dial(grpcListen, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
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
	err := test.clientConn.Close()
	if err != nil {
		fmt.Errorf("error on close connection: %v", err)
	}
}

func FeatureContextGrpc(s *godog.Suite) {
	testGrpc := new(rotationGrpcTest)

	s.BeforeScenario(testGrpc.connect)

	s.AfterScenario(testGrpc.close)
}
