SHELL:="/bin/bash"
MAIN_VERSION:=$(shell git describe --always --abbrev=0 --tags || echo "0.1")
PLATFORM:=$(shell go env GOOS)
ARCH:=$(shell go env GOARCH)
GOPATH:=$(shell go env GOPATH)
GOBIN:=$(GOPATH)/bin
LDFLAGS:="-X github.com/jorgechato/acictl/utils.VERSION=${MAIN_VERSION}"


.PHONY: release

default: test build

test: linter cov unit

install_dep:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

install_dev:
	go get -u gopkg.in/alecthomas/gometalinter.v2
	gometalinter.v2 --install

install_unit:
	go get -u github.com/jstemmer/go-junit-report

build:
	go build -ldflags ${LDFLAGS} -a -o acictl main.go

run:
	go run -ldflags ${LDFLAGS} main.go

linter:
	gometalinter.v2 --checkstyle > report.xml

cov:
	go test -coverprofile=coverage.out ./...

unit:
	go test -v ./... | go-junit-report > test.xml

clean:
	rm -rf acictl *.out *.xml release

release:
	./lazy/build.sh ${LDFLAGS}
