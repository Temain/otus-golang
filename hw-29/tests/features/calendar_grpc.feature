Feature: gRPC requests handling
	As API client of calendar service
	In order to understand that the calendar working
	I want to receive simple request

	Scenario: Calendar gRPC API service add method is available
		When I send request to add method
		Then The add request response code should be 0 (ok)
		And The add request response should be with value true

	Scenario: Calendar gRPC API service add method can't add duplicate events
		When I send request to add method with existing event
		Then The add request response code should be 13 (internal error)
		And The add request response should be with value false

	Scenario: Calendar gRPC API service list method is available
		When I send request to list method
		Then The list request response code should be 0 (ok)
		And The list request response should be non empty

	Scenario: Calendar gRPC API service search method is available
		When I send request to search method
		Then The search request response code should be 0 (ok)
		And The search request response should contains 1 event

	Scenario: Calendar gRPC API service update method is available
		When I send request to update method
		Then The update request response code should be 0 (ok)
		And The update request response should be with value true

	Scenario: Calendar gRPC API service update method can't update not existing event
		When I send request to update method with not existing event
		Then The update request response code should be 13 (internal error)
		And The update request response should be with value false

	Scenario: Calendar gRPC API service delete method is available
		When I send request to delete method
		Then The delete request response code should be 0 (ok)
		And The delete request response should be with value true

	Scenario: Calendar gRPC API service delete method can't delete not existing event
		When I send request to delete method with not existing event
		Then The delete request response code should be 13 (internal error)
		And The delete request response should be with value false
