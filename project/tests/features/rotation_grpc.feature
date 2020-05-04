Feature: gRPC requests handling
	As API client of banner rotation service
	In order to understand that the service working
	I want to receive simple requests

	Scenario: Calendar gRPC API service add method is available
		When I send request to add banner method
		Then The add banner request response code should be 0 (ok)
		And The add banner request response should be with value true