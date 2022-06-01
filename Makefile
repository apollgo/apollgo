.DEFAULT_GOAL := build

.PHONY: clean test lint build update

clean:
	rm aenema

test:
	go test ./...

lint:
	golangci-lint run

build:
	go build -o aenema ./cmd/aenema

update:
	go get -u ./cmd/aenema
