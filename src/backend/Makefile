all: clean test documentation docker
.PHONY: all

build:
	go mod download
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
	go build -installsuffix cgo -o app ./main.go

documentation:
	go install github.com/swaggo/swag/cmd/swag
	swag init -g ./main.go

test:
	go test -short -cover -v -coverprofile=coverage.out -covermode=atomic ./...

clean:
	go clean

docker:
	docker run --rm -v ~/.netrc:/root/.netrc -v $(shell pwd):/src -w /src golang:1.16 make build
	docker build -t zooservice .
