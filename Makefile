SHELL="/bin/bash"
PLATFORM=$(shell go env GOOS)
ARCH=$(shell go env GOARCH)
GOPATH=$(shell go env GOPATH)
GOBIN=$(GOPATH)/bin

default: build test

install_dep:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

get_dep:
	go get -t -v ./...

build:
	go install

test:
	./lazy/coverage.sh
