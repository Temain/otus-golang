package main

import (
	"context"

	"github.com/cucumber/godog"
	"google.golang.org/grpc"

	"github.com/Temain/otus-golang/project/pkg/proto"
)

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

func (test *rotationGrpcTest) iCallAddMethod() error {
	return godog.ErrPending
}

func (test *rotationGrpcTest) methodShouldReturnSuccessResult() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	testGrpc := new(rotationGrpcTest)

	s.Step(`^I call add method$`, testGrpc.iCallAddMethod)
	s.Step(`^Method should return success result$`, testGrpc.methodShouldReturnSuccessResult)
}
