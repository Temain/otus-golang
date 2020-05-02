package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cucumber/godog"
	"google.golang.org/grpc"
	// event "github.com/Temain/otus-golang/hw-29/pkg/proto"
)

type calendarTest struct {
	ctx                context.Context
	conn               *grpc.ClientConn
	responseStatusCode int
	responseBody       []byte
}

func (test *calendarTest) dialGrpc() (err error) {
	/*test.ctx, _ = context.WithTimeout(context.Background(), 5*time.Minute)
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	conn := event.NewEventServiceClient(cc)*/
	return err
}

func (test *calendarTest) iSendRequestTo(httpMethod, addr string) (err error) {
	var r *http.Response

	switch httpMethod {
	case http.MethodGet:
		r, err = http.Get(addr)
	default:
		err = fmt.Errorf("unknown method: %s", httpMethod)
	}

	if err != nil {
		return
	}
	test.responseStatusCode = r.StatusCode
	test.responseBody, err = ioutil.ReadAll(r.Body)
	return
}

func (test *calendarTest) theResponseCodeShouldBe(code int) (err error) {
	if test.responseStatusCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", test.responseStatusCode, code)
	}
	return nil
}

func (test *calendarTest) theResponseShouldMatchText(text string) (err error) {
	if string(test.responseBody) != text {
		return fmt.Errorf("unexpected text: %s != %s", test.responseBody, text)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	test := new(calendarTest)

	// s.BeforeScenario(test.dialGrpc)

	s.Step(`^I send "([^"]*)" request to "([^"]*)"$`, test.iSendRequestTo)
	s.Step(`^The response code should be (\d+)$`, test.theResponseCodeShouldBe)
	s.Step(`^The response should match text "([^"]*)"$`, test.theResponseShouldMatchText)
}
