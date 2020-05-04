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
	"google.golang.org/grpc/status"

	"github.com/Temain/otus-golang/project/pkg/proto"
)

var grpcListen = os.Getenv("TESTS_GRPC_API")

func init() {
	if grpcListen == "" {
		grpcListen = "rotation_grpc_api:50052"
	}
}

type rotationGrpcTest struct {
	ctx             context.Context
	clientConn      *grpc.ClientConn
	client          proto.RotationServiceClient
	addBannerResult bool
	addBannerCode   int
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

func getStatusCode(err error) (code int) {
	statusErr, ok := status.FromError(err)
	if ok {
		code = int(statusErr.Code())
	} else {
		code = -1
	}
	return code
}

func (test *rotationGrpcTest) iSendRequestToAddBannerMethod() error {
	request := &proto.AddBannerRequest{}
	response, err := test.client.AddBanner(test.ctx, request)
	if err != nil {
		test.addBannerCode = getStatusCode(err)
		return fmt.Errorf("error on add banner to rotation: %v", err)
	}

	test.addBannerResult = response.Success

	return nil
}

func (test *rotationGrpcTest) theAddBannerRequestResponseCodeShouldBeOk(code int) error {
	if test.addBannerCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", test.addBannerCode, code)
	}
	return nil
}

func (test *rotationGrpcTest) theAddBannerRequestResponseShouldBeWithValueTrue() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	testGrpc := new(rotationGrpcTest)

	s.BeforeScenario(testGrpc.connect)

	s.Step(`^I send request to add banner method$`, testGrpc.iSendRequestToAddBannerMethod)
	s.Step(`^The add banner request response code should be (\d+) \(ok\)$`, testGrpc.theAddBannerRequestResponseCodeShouldBeOk)
	s.Step(`^The add banner request response should be with value true$`, testGrpc.theAddBannerRequestResponseShouldBeWithValueTrue)

	s.AfterScenario(testGrpc.close)
}
