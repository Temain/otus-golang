Feature: gRPC requests handling
	As API client of banner rotation service
	In order to understand that the service working
	I want to receive simple requests

	Scenario: Calendar gRPC API service add banner method is available
		When I send request to add banner method
		Then The add banner request response code should be 0 (ok)
		And The add banner request response should be with value true

	Scenario: Calendar gRPC API service delete banner method is available
		When I send request to delete banner method
		Then The delete banner request response code should be 0 (ok)
		And The delete banner request response should be with value true

	Scenario: Calendar gRPC API service click banner method is available
		When I send request to click banner method
		Then The click banner request response code should be 0 (ok)
		And The click banner request response should be with value true

	Scenario: Calendar gRPC API service get banner method is available
		When I send request to get banner method
		Then The get banner request response code should be 0 (ok)
		And The get banner request response should be not 0

