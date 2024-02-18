.DEFAULT_GOAL := build

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet: fmt
	go vet ./...

.PHONY: build
build: vet
	go build

.PHONY: clear
clean:
	go clean
