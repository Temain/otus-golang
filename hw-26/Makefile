PATH_API=./cmd/calendar_api
PATH_SCHEDULER=./cmd/calendar_scheduler
PATH_SENDER=./cmd/calendar_sender

build:
	go build $(PATH_API)
	go build $(PATH_SCHEDULER)
	go build $(PATH_SENDER)

run-grpc-api:
	go run $(PATH_API) grpc_server --config=configs/config_api.json

run-scheduler:
	go run $(PATH_SCHEDULER) --config=configs/config_scheduler.json

run-sender:
	go run $(PATH_SENDER) --config=configs/config_sender.json