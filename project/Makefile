PATH_API=./cmd/rotation_api

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: build
build:
	go build $(PATH_API)

.PHONY: run-api
run-api:
	go run $(PATH_API) --config=configs/config.json

.PHONY: run
run:
	docker-compose up -d --build

.PHONY: stop
stop:
	docker-compose down

.PHONY: restart
restart: stop run

.PHONY: test
test:
	set -e ;\
	docker-compose -f docker-compose.test.yml up --build -d ;\
	test_status_code=0 ;\
	docker-compose -f docker-compose.test.yml run integration_tests go test || test_status_code=$$? ;\
	docker-compose -f docker-compose.test.yml down ;\
	exit $$test_status_code ;\