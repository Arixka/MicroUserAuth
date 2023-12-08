build:
	go build -o /app/microUserAuth cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test -v ./test/...

build-docker: build
	docker build . -t api-rest

run-docker: build-docker
	docker run -p 3000:3000 api-rest+

down-docker:
	docker-compose down --remove-orphans

docker-logs:
	docker logs [CONTAINER_ID_OR_NAME]
