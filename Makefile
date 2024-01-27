build:
	go build -o /app/microUserAuth cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test -v ./test/...

build-docker:
	docker-compose build

run-docker:
	docker-compose up

down-docker:
	docker-compose down --remove-orphans

docker-logs:
	docker logs [CONTAINER_ID_OR_NAME]
