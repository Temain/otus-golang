package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cucumber/godog"
)

type calendarHttpTest struct {
	responseStatusCode int
	responseBody       []byte
}

func (test *calendarHttpTest) iSendRequestTo(httpMethod, addr string) (err error) {
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

func (test *calendarHttpTest) theResponseCodeShouldBe(code int) (err error) {
	if test.responseStatusCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", test.responseStatusCode, code)
	}
	return nil
}

func (test *calendarHttpTest) theResponseShouldMatchText(text string) (err error) {
	if string(test.responseBody) != text {
		return fmt.Errorf("unexpected text: %s != %s", test.responseBody, text)
	}
	return nil
}

func FeatureContextHttp(s *godog.Suite) {
	test := new(calendarHttpTest)

	s.Step(`^I send "([^"]*)" request to "([^"]*)"$`, test.iSendRequestTo)
	s.Step(`^The response code should be (\d+)$`, test.theResponseCodeShouldBe)
	s.Step(`^The response should match text "([^"]*)"$`, test.theResponseShouldMatchText)
}
