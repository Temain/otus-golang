PATH_API=./cmd/rotation_api

build:
	go build $(PATH_API)

run-api:
	go run $(PATH_API) --config=configs/config.json

up:
	docker-compose up -d --build

down:
	docker-compose down

restart: down up

test:
	set -e ;\
	docker-compose -f docker-compose.test.yml up --build -d ;\
	test_status_code=0 ;\
	docker-compose -f docker-compose.test.yml run integration_tests go test || test_status_code=$$? ;\
	docker-compose -f docker-compose.test.yml down ;\
	exit $$test_status_code ;\
