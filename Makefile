.DEFAULT_GOAL := build

.PHONY: lint test build clean update

lint:
	golangci-lint run

test:
	go test ./...

build:
	go build -o aenema ./cmd/aenema

update:
	go get -u ./cmd/aenema

clean:
	rm aenema
