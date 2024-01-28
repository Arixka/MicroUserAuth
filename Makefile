build:
	go build -o ./out/microUserAuth ./cmd/api/main.go

run:
	go build -o ./out/microUserAuth ./cmd/api/main.go
	copy .env.local .env
	.\out\microUserAuth

run-local:
	copy .env.local .env
	go run ./cmd/api/main.go

test:
	go test -v ./test/...

build-docker:
	docker-compose build

run-docker:
	cp .env.docker .env
	docker-compose up

down-docker:
	docker-compose down --remove-orphans

docker-logs:
	docker logs [CONTAINER_ID_OR_NAME]

run-db:
	docker-compose up db