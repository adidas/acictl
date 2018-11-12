SHELL:="/bin/bash"
MAIN_VERSION:=$(shell git describe --always --abbrev=0 --tags || echo "0.1")
PLATFORM:=$(shell go env GOOS)
ARCH:=$(shell go env GOARCH)
GOPATH:=$(shell go env GOPATH)
GOBIN:=$(GOPATH)/bin
LDFLAGS:="-X github.com/adidas/acictl/utils.VERSION=${MAIN_VERSION}"
M = $(shell printf "\033[34;1m▶\033[0m")


.PHONY: build clean dep schema help default server setup vendor release test

default: help

vendor: ; $(info $(M) Download dev dependencies...)
	go get github.com/oxequa/realize
	go get -u github.com/tebeka/go2xunit

setup: ; $(info $(M) Fetching github.com/golang/dep...)
	go get github.com/golang/dep/cmd/dep

dep: setup ; $(info $(M) Ensuring vendored dependencies are up-to-date...) ## Download dependencies
	dep ensure -v -vendor-only

build: dep ; $(info $(M) Building project...) ## Build server binary
	go build -ldflags ${LDFLAGS} -a -o acictl main.go

run:
	go run -ldflags ${LDFLAGS} main.go

clean: ; $(info $(M) Removing generated files... ) ## Clean schema bind data
	rm -rf acictl *.out *.xml release

release: ; $(info $(M) Release project...) ## Release binary
	./lazy/build.sh ${LDFLAGS}

test: vendor ; $(info $(M) Run tests...) ## Start unit tests
	go test -race -cover -v -coverprofile=coverage.out ./... > test.output
	cat test.output | go2xunit -output tests.xml

server: vendor ; $(info $(M) Starting development server...) ## Start development server
	realize start

unit: vendor ; $(info $(M) Run unit tests...) ## Start unit tests
	go list $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=60s -parallel=4

fmtcheck:
	@sh -c "'$(CURDIR)/lazy/gofmtcheck.sh'"

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
