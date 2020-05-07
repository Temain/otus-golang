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
	ctx                context.Context
	clientConn         *grpc.ClientConn
	client             proto.RotationServiceClient
	addBannerResult    bool
	addBannerCode      int
	deleteBannerResult bool
	deleteBannerCode   int
	clickBannerResult  bool
	clickBannerCode    int
	getBannerResult    int64
	getBannerCode      int
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
	request := &proto.AddBannerRequest{BannerId: 4, SlotId: 1}
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
	if !test.addBannerResult {
		return fmt.Errorf("banner not added to rotation")
	}
	return nil
}

func (test *rotationGrpcTest) iSendRequestToDeleteBannerMethod() error {
	request := &proto.DeleteBannerRequest{BannerId: 4, SlotId: 1}
	response, err := test.client.DeleteBanner(test.ctx, request)
	if err != nil {
		test.deleteBannerCode = getStatusCode(err)
		return fmt.Errorf("error on delete banner from rotation: %v", err)
	}

	test.deleteBannerResult = response.Success

	return nil
}

func (test *rotationGrpcTest) theDeleteBannerRequestResponseCodeShouldBeOk(code int) error {
	if test.deleteBannerCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", test.deleteBannerCode, code)
	}
	return nil
}

func (test *rotationGrpcTest) theDeleteBannerRequestResponseShouldBeWithValueTrue() error {
	if !test.deleteBannerResult {
		return fmt.Errorf("banner not deleted from rotation")
	}
	return nil
}

func (test *rotationGrpcTest) iSendRequestToClickBannerMethod() error {
	request := &proto.ClickBannerRequest{BannerId: 1, SlotId: 1, GroupId: 1}
	response, err := test.client.ClickBanner(test.ctx, request)
	if err != nil {
		test.deleteBannerCode = getStatusCode(err)
		return fmt.Errorf("error on click banner: %v", err)
	}

	test.clickBannerResult = response.Success

	return nil
}

func (test *rotationGrpcTest) theClickBannerRequestResponseCodeShouldBeOk(code int) error {
	if test.clickBannerCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", test.clickBannerCode, code)
	}
	return nil
}

func (test *rotationGrpcTest) theClickBannerRequestResponseShouldBeWithValueTrue() error {
	if !test.clickBannerResult {
		return fmt.Errorf("error on click banner")
	}
	return nil
}

func (test *rotationGrpcTest) iSendRequestToGetBannerMethod() error {
	request := &proto.GetBannerRequest{SlotId: 1, GroupId: 1}
	response, err := test.client.GetBanner(test.ctx, request)
	if err != nil {
		test.deleteBannerCode = getStatusCode(err)
		return fmt.Errorf("error on get banner: %v", err)
	}

	test.getBannerResult = response.BannerId

	return nil
}

func (test *rotationGrpcTest) theGetBannerRequestResponseCodeShouldBeOk(code int) error {
	if test.getBannerCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", test.getBannerCode, code)
	}
	return nil
}

func (test *rotationGrpcTest) theGetBannerRequestResponseShouldBeNot(arg1 int) error {
	if test.getBannerResult == 0 {
		return fmt.Errorf("wrong get banner result, must be non zero")
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	testGrpc := new(rotationGrpcTest)

	s.BeforeScenario(testGrpc.connect)

	s.Step(`^I send request to add banner method$`, testGrpc.iSendRequestToAddBannerMethod)
	s.Step(`^The add banner request response code should be (\d+) \(ok\)$`, testGrpc.theAddBannerRequestResponseCodeShouldBeOk)
	s.Step(`^The add banner request response should be with value true$`, testGrpc.theAddBannerRequestResponseShouldBeWithValueTrue)

	s.Step(`^I send request to click banner method$`, testGrpc.iSendRequestToClickBannerMethod)
	s.Step(`^The click banner request response code should be (\d+) \(ok\)$`, testGrpc.theClickBannerRequestResponseCodeShouldBeOk)
	s.Step(`^The click banner request response should be with value true$`, testGrpc.theClickBannerRequestResponseShouldBeWithValueTrue)

	s.Step(`^I send request to delete banner method$`, testGrpc.iSendRequestToDeleteBannerMethod)
	s.Step(`^The delete banner request response code should be (\d+) \(ok\)$`, testGrpc.theDeleteBannerRequestResponseCodeShouldBeOk)
	s.Step(`^The delete banner request response should be with value true$`, testGrpc.theDeleteBannerRequestResponseShouldBeWithValueTrue)

	s.Step(`^I send request to get banner method$`, testGrpc.iSendRequestToGetBannerMethod)
	s.Step(`^The get banner request response code should be (\d+) \(ok\)$`, testGrpc.theGetBannerRequestResponseCodeShouldBeOk)
	s.Step(`^The get banner request response should be not (\d+)$`, testGrpc.theGetBannerRequestResponseShouldBeNot)

	s.AfterScenario(testGrpc.close)
}
