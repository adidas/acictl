SHELL="/bin/bash"
PLATFORM=$(shell go env GOOS)
ARCH=$(shell go env GOARCH)
GOPATH=$(shell go env GOPATH)
GOBIN=$(GOPATH)/bin

default: test build

test: unit cov linter

install_dep:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

install_dev:
	go get -u gopkg.in/alecthomas/gometalinter.v2
	gometalinter.v2 --install

install_unit:
	go get -u github.com/jstemmer/go-junit-report

build:
	go install

run:
	go run ${LDFLAGS} manage.go

linter:
	gometalinter.v2 --checkstyle > report.xml

cov:
	go test -coverprofile=coverage.out ./...

unit:
	go test -v ./... | go-junit-report > test.xml

clean:
	rm -rf server *.out *.xml
