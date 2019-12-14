.PHONY: build test

build:
	@go build -o treego

test:
	@go test -v ./...
