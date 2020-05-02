Feature: gRPC requests handling
	As API client of calendar service
	In order to understand that the calendar working
	I want to receive simple request

	Scenario: Calendar gRPC API service list method is available
		When I call list method
		Then The result should be non empty

	Scenario: Calendar gRPC API service search method is available
		When I call search method
		Then Method should return 1 event
