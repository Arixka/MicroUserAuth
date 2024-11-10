
OS := $(shell uname -s)

ifeq ($(OS), Linux)
    COPY = cp
else
    COPY = copy
endif

build:
	go build -o ./out/microUserAuth.exe ./cmd/api/main.go

run-db:
	docker-compose up db

run:
	go build -o ./out/microUserAuth ./cmd/api/main.go
	$(COPY) .env.local .env
	./out/microUserAuth

run-local:
	$(COPY) .env.local .env
	go run ./cmd/api/main.go

test:
	go test -v ./test/...

build-docker:
	docker-compose build

run-docker:
	$(COPY) .env.docker .env
	docker-compose up -d

down-docker:
	docker-compose down --remove-orphans

docker-logs:
	docker logs [CONTAINER_ID_OR_NAME]
