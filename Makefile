
OS := $(shell uname -s)

ifeq ($(OS), Linux)
    COPY = cp
else
    COPY = copy
endif

build:
	go build -o ./out/microUserAuth.exe ./cmd/api/main.go

run-db:
	docker start docker_db || docker-compose -f ../../Docker_DB/docker-compose-db.yml up -d

run: build
	./out/microUserAuth.exe

dev:# Hot reload (desarrollo)
	air

run-local: build # Sin reload
	go run ./cmd/api/main.go

test:
	go test -v ./test/...

build-docker:
	docker-compose build

run-docker:
	docker-compose up -d

down-docker:
	docker-compose down --remove-orphans

docker-logs:
	docker logs [CONTAINER_ID_OR_NAME]
