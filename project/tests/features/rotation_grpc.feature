Feature: gRPC requests handling
	As API client of banner rotation service
	In order to understand that the service working
	I want to receive simple requests

	Scenario: Rotation gRPC API service add method is available
		When I call add method
		Then Method should return success result