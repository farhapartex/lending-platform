SERVICE := core-service
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
LDFLAGS := -s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT)

.PHONY: tidy fmt vet build run test docker-build

tidy:
	go mod tidy

fmt:
	go fmt ./...

vet:
	go vet ./...

build:
	CGO_ENABLED=0 go build -trimpath -ldflags="$(LDFLAGS)" -o bin/$(SERVICE) ./cmd/api

run:
	go run ./cmd/api

test:
	go test ./... -race -count=1

docker-build:
	docker build --build-arg VERSION=$(VERSION) --build-arg COMMIT=$(COMMIT) -t $(SERVICE):$(VERSION) .
