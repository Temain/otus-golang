package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	event "github.com/Temain/otus-golang/hw-29/pkg/proto"
)

var grpcListen = os.Getenv("TESTS_GRPC_API")

func init() {
	if grpcListen == "" {
		grpcListen = "grpc_api:50051"
	}
}

type calendarGrpcTest struct {
	ctx                 context.Context
	clientConn          *grpc.ClientConn
	client              event.EventServiceClient
	sampleEvent         *event.EventMessage
	listResult          []event.EventMessage
	listCode            int
	searchResult        *event.EventMessage
	searchCode          int
	addResult           bool
	addCode             int
	addDuplicateResult  bool
	addDuplicateCode    int
	updateResult        bool
	updateCode          int
	updateNotExists     bool
	updateNotExistsCode int
	deleteResult        bool
	deleteCode          int
	deleteNotExists     bool
	deleteNotExistsCode int
}

func (test *calendarGrpcTest) connect(*messages.Pickle) {
	test.ctx, _ = context.WithTimeout(context.Background(), 5*time.Minute)
	cc, err := grpc.Dial(grpcListen, grpc.WithInsecure())
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

func getStatusCode(err error) (code int) {
	statusErr, ok := status.FromError(err)
	if ok {
		code = int(statusErr.Code())
	} else {
		code = -1
	}
	return code
}

func (test *calendarGrpcTest) iSendRequestToAddMethod() error {
	request := &event.AddRequest{Event: test.sampleEvent}
	response, err := test.client.Add(test.ctx, request)
	if err != nil {
		test.addCode = getStatusCode(err)
		return fmt.Errorf("error on add event: %v", err)
	}

	test.addResult = response.Success

	return nil
}

func (test *calendarGrpcTest) theAddRequestResponseCodeShouldBeOk(code int) error {
	if test.addCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", test.addDuplicateCode, code)
	}
	return nil
}

func (test *calendarGrpcTest) theAddRequestResponseShouldBeWithValueTrue() error {
	if !test.addResult {
		return fmt.Errorf("new event not added")
	}
	return nil
}

func (test *calendarGrpcTest) iSendRequestToAddMethodWithExistingEvent() error {
	request := &event.AddRequest{Event: test.sampleEvent}
	_, err := test.client.Add(test.ctx, request)
	if err != nil {
		test.addDuplicateCode = getStatusCode(err)
		test.addDuplicateResult = false
	}
	return nil
}

func (test *calendarGrpcTest) theAddRequestResponseCodeShouldBeInternalError(code int) error {
	if test.addDuplicateCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", test.addDuplicateCode, code)
	}
	return nil
}

func (test *calendarGrpcTest) theAddRequestResponseShouldBeWithValueFalse() error {
	if test.addDuplicateResult {
		return fmt.Errorf("duplicate event added")
	}
	return nil
}

func (test *calendarGrpcTest) iSendRequestToListMethod() error {
	stream, err := test.client.List(test.ctx, &event.ListRequest{})
	if err != nil {
		test.listCode = getStatusCode(err)
		err = fmt.Errorf("error on list events: %v", err)
	}

	if stream == nil {
		err = fmt.Errorf("wrong result on list of events: %v", err)
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
		test.listResult = append(test.listResult, *msg)
	}

	return err
}

func (test *calendarGrpcTest) theListRequestResponseCodeShouldBeOk(code int) error {
	if test.listCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", test.addDuplicateCode, code)
	}
	return nil
}

func (test *calendarGrpcTest) theListRequestResponseShouldBeNonEmpty() error {
	if len(test.listResult) == 0 {
		return fmt.Errorf("result of list method is empty")
	}
	return nil
}

func (test *calendarGrpcTest) iSendRequestToSearchMethod() error {
	created := test.sampleEvent.Created
	response, err := test.client.Search(test.ctx, &event.SearchRequest{Date: created})
	if err != nil {
		test.searchCode = getStatusCode(err)
		return fmt.Errorf("error on search event: %v", err)
	}

	test.searchResult = response.Event

	return nil
}

func (test *calendarGrpcTest) theSearchRequestResponseCodeShouldBeOk(code int) error {
	if test.searchCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", test.addDuplicateCode, code)
	}
	return nil
}

func (test *calendarGrpcTest) theSearchRequestResponseShouldContainsEvent() error {
	if test.searchResult == nil {
		return fmt.Errorf("result of search method is empty")
	}
	return nil
}

func (test *calendarGrpcTest) iSendRequestToUpdateMethod() error {
	request := &event.UpdateRequest{Event: test.sampleEvent}
	response, err := test.client.Update(test.ctx, request)
	if err != nil {
		test.updateCode = getStatusCode(err)
		return fmt.Errorf("error on update event: %v", err)
	}

	test.updateResult = response.Success

	return nil
}

func (test *calendarGrpcTest) theUpdateRequestResponseCodeShouldBeOk(code int) error {
	if test.updateCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", test.addDuplicateCode, code)
	}
	return nil
}

func (test *calendarGrpcTest) theUpdateRequestResponseShouldBeWithValueTrue() error {
	if !test.updateResult {
		return fmt.Errorf("event not updated")
	}
	return nil
}

func (test *calendarGrpcTest) iSendRequestToUpdateMethodWithNotExistingEvent() error {
	created, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		return fmt.Errorf("wrong event date: %v", err)
	}
	notExists := &event.EventMessage{
		Id:          2,
		Title:       "Not exists event title",
		Description: "Not exists event description",
		Created:     created,
	}
	request := &event.UpdateRequest{Event: notExists}
	_, err = test.client.Update(test.ctx, request)
	if err != nil {
		test.updateNotExistsCode = getStatusCode(err)
		test.updateNotExists = false
	}

	return nil
}

func (test *calendarGrpcTest) theUpdateRequestResponseCodeShouldBeInternalError(code int) error {
	if test.updateNotExistsCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", test.addDuplicateCode, code)
	}
	return nil
}

func (test *calendarGrpcTest) theUpdateRequestResponseShouldBeWithValueFalse() error {
	if test.updateNotExists {
		return fmt.Errorf("not existing event updated")
	}
	return nil
}

func (test *calendarGrpcTest) iSendRequestToDeleteMethod() error {
	request := &event.DeleteRequest{Id: 1}
	response, err := test.client.Delete(test.ctx, request)
	if err != nil {
		test.deleteCode = getStatusCode(err)
		return fmt.Errorf("error on delete event: %v", err)
	}

	test.deleteResult = response.Success

	return nil
}

func (test *calendarGrpcTest) theDeleteRequestResponseCodeShouldBeOk(code int) error {
	if test.deleteCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", test.addDuplicateCode, code)
	}
	return nil
}

func (test *calendarGrpcTest) theDeleteRequestResponseShouldBeWithValueTrue() error {
	if !test.deleteResult {
		return fmt.Errorf("event not deleted")
	}
	return nil
}

func (test *calendarGrpcTest) iSendRequestToDeleteMethodWithNotExistingEvent() error {
	request := &event.DeleteRequest{Id: 2}
	_, err := test.client.Delete(test.ctx, request)
	if err != nil {
		test.deleteNotExistsCode = getStatusCode(err)
		test.deleteNotExists = false
	}

	return nil
}

func (test *calendarGrpcTest) theDeleteRequestResponseCodeShouldBeInternalError(code int) error {
	if test.deleteNotExistsCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", test.addDuplicateCode, code)
	}
	return nil
}

func (test *calendarGrpcTest) theDeleteRequestResponseShouldBeWithValueFalse() error {
	if test.deleteNotExists {
		return fmt.Errorf("not existing event deleted")
	}
	return nil
}

func FeatureContextGrpc(s *godog.Suite) {
	testGrpc := new(calendarGrpcTest)

	s.BeforeScenario(testGrpc.connect)

	s.Step(`^I send request to add method$`, testGrpc.iSendRequestToAddMethod)
	s.Step(`^The add request response code should be (\d+) \(ok\)$`, testGrpc.theAddRequestResponseCodeShouldBeOk)
	s.Step(`^The add request response should be with value true$`, testGrpc.theAddRequestResponseShouldBeWithValueTrue)

	s.Step(`^I send request to add method with existing event$`, testGrpc.iSendRequestToAddMethodWithExistingEvent)
	s.Step(`^The add request response code should be (\d+) \(internal error\)$`, testGrpc.theAddRequestResponseCodeShouldBeInternalError)
	s.Step(`^The add request response should be with value false$`, testGrpc.theAddRequestResponseShouldBeWithValueFalse)

	s.Step(`^I send request to list method$`, testGrpc.iSendRequestToListMethod)
	s.Step(`^The list request response code should be (\d+) \(ok\)$`, testGrpc.theListRequestResponseCodeShouldBeOk)
	s.Step(`^The list request response should be non empty$`, testGrpc.theListRequestResponseShouldBeNonEmpty)

	s.Step(`^I send request to search method$`, testGrpc.iSendRequestToSearchMethod)
	s.Step(`^The search request response code should be (\d+) \(ok\)$`, testGrpc.theSearchRequestResponseCodeShouldBeOk)
	s.Step(`^The search request response should contains (\d+) event$`, testGrpc.theSearchRequestResponseShouldContainsEvent)

	s.Step(`^I send request to update method$`, testGrpc.iSendRequestToUpdateMethod)
	s.Step(`^The update request response code should be (\d+) \(ok\)$`, testGrpc.theUpdateRequestResponseCodeShouldBeOk)
	s.Step(`^The update request response should be with value true$`, testGrpc.theUpdateRequestResponseShouldBeWithValueTrue)

	s.Step(`^I send request to update method with not existing event$`, testGrpc.iSendRequestToUpdateMethodWithNotExistingEvent)
	s.Step(`^The update request response code should be (\d+) \(internal error\)$`, testGrpc.theUpdateRequestResponseCodeShouldBeInternalError)
	s.Step(`^The update request response should be with value false$`, testGrpc.theUpdateRequestResponseShouldBeWithValueFalse)

	s.Step(`^I send request to delete method$`, testGrpc.iSendRequestToDeleteMethod)
	s.Step(`^The delete request response code should be (\d+) \(ok\)$`, testGrpc.theDeleteRequestResponseCodeShouldBeOk)
	s.Step(`^The delete request response should be with value true$`, testGrpc.theDeleteRequestResponseShouldBeWithValueTrue)

	s.Step(`^I send request to delete method with not existing event$`, testGrpc.iSendRequestToDeleteMethodWithNotExistingEvent)
	s.Step(`^The delete request response code should be (\d+) \(internal error\)$`, testGrpc.theDeleteRequestResponseCodeShouldBeInternalError)
	s.Step(`^The delete request response should be with value false$`, testGrpc.theDeleteRequestResponseShouldBeWithValueFalse)

	s.AfterScenario(testGrpc.close)
}
