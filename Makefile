.DEFAULT_GOAL := build

export GO111MODULE=on
export CGO_ENABLED=0
export BINARY=apollgo
export BUILD_DATE=$(shell LANG=en_us_88591 date)
export BUILD=$(shell git rev-parse HEAD 2> /dev/null || echo "undefined")
export VERSION=$(shell git describe --abbrev=0 --tags 2> /dev/null || echo "0.1.0")
export LD_FLAGS=-ldflags "-X 'github.com/apollgo/apollgo.Version=$(VERSION)' -s -w"

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build
	go build $(LD_FLAGS) -o $(BINARY)

.PHONY: tmp
tmp: ## Build and output the binary in /tmp
	go build $(LD_FLAGS) -o /tmp/$(BINARY)

.PHONY: packed
packed: ## Build a packed version of the binary
	go build $(LD_FLAGS) -o $(BINARY)
	upx --best --lzma $(BINARY)

.PHONY: docker
docker: ## Build the docker image with packed binaries
	docker build -t $(BINARY):latest -t $(BINARY):$(BUILD) -f Dockerfile .

.PHONY: lint
lint: ## Runs the linter
	$(GOPATH)/bin/golangci-lint run --exclude-use-default=false

.PHONY: test
test: ## Run the test suite
	CGO_ENABLED=1 go test -race -coverprofile="coverage.txt" ./...

.PHONY: clean
clean: ## Remove the binary
	if [ -f $(BINARY) ] ; then rm $(BINARY) ; fi
	if [ -f coverage.txt ] ; then rm coverage.txt ; fi
