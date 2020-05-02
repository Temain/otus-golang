# http://localhost:8888/

Feature: gRPC requests handling
	As API client of calendar service
	In order to understand that the calendar working
	I want to receive simple request

	Scenario: Calendar GRPC API service is available
		When I call list method
		Then The result should be non empty
