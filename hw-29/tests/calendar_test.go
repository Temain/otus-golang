package main

import "github.com/cucumber/godog"

func iSendRequestTo(arg1, arg2 string) error {
	return godog.ErrPending
}

func theResponseCodeShouldBe(arg1 int) error {
	return godog.ErrPending
}

func theResponseShouldMatchText(arg1 string) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I send "([^"]*)" request to "([^"]*)"$`, iSendRequestTo)
	s.Step(`^The response code should be (\d+)$`, theResponseCodeShouldBe)
	s.Step(`^The response should match text "([^"]*)"$`, theResponseShouldMatchText)
}
