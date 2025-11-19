	.PHONY: help build run test clean fmt lint

help:
	@echo "Available commands:"
	@echo "  make build    - Build the server binary"
	@echo "  make run      - Run the server"
	@echo "  make test     - Run tests"
	@echo "  make clean    - Remove build artifacts"
	@echo "  make fmt      - Format code"
	@echo "  make lint     - Run linter"

build:
	go build -o bin/server ./cmd/server

run:
	go run ./cmd/server

test:
	go test -v ./...

clean:
	rm -rf bin/

fmt:
	go fmt ./...

lint:
	golangci-lint run ./...
