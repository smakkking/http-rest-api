:PHONY build
build:
	go build -v ./cmd/apiserver

:PHONY test
test:
	go test -v -race -timeout 3s ./...

.DEFAULT_GOAL := build