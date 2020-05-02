# http://localhost:8888/

Feature: HTTP requests handling
	As API client of calendar service
	In order to understand that the calendar working
	I want to receive simple request

	Scenario: Calendar API service is available
		When I send "GET" request to "http://api:8888/hello"
		Then The response code should be 200
		And The response should match text "hello"
