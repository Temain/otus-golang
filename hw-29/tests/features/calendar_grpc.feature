Feature: gRPC requests handling
	As API client of calendar service
	In order to understand that the calendar working
	I want to receive simple request

	Scenario: Calendar gRPC API service add method is available
		When I call add method
		Then Method should return success result

	Scenario: Calendar gRPC API service add method can't add duplicate events
		When I call add method with existing event
		Then Method should return fail result

	Scenario: Calendar gRPC API service list method is available
		When I call list method
		Then The result should be non empty

	Scenario: Calendar gRPC API service search method is available
		When I call search method
		Then Method should return 1 event

	Scenario: Calendar gRPC API service update method is available
		When I call update method
		Then Method should return success result

	Scenario: Calendar gRPC API service update method can't update not existing event
		When I call update method with not existing event
		Then Method should return fail result

	Scenario: Calendar gRPC API service delete method is available
		When I call delete method
		Then Method should return success result

	Scenario: Calendar gRPC API service delete method can't delete not existing event
		When I call delete method with not existing event
		Then Method should return fail result
