.PHONY: build clean test lint tidy run install release

BINARY=goforge
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS=-ldflags "-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)"

build:
	go build $(LDFLAGS) -o $(BINARY) ./cmd/goforge

install:
	go install $(LDFLAGS) ./cmd/goforge

test:
	go test ./... -v -cover -count=1

test-race:
	go test ./... -race -count=1

test-cover:
	go test ./... -coverprofile=coverage.out -count=1
	go tool cover -html=coverage.out -o coverage.html

lint:
	golangci-lint run ./...

tidy:
	go mod tidy
	go mod verify

clean:
	rm -f $(BINARY)
	rm -f coverage.out coverage.html
	rm -rf release/

run:
	go run ./cmd/goforge

release:
	goreleaser release --clean
