SHELL:="/bin/bash"
MAIN_VERSION:=$(shell git describe --always --abbrev=0 --tags || echo "0.1")
PLATFORM:=$(shell go env GOOS)
ARCH:=$(shell go env GOARCH)
GOPATH:=$(shell go env GOPATH)
GOBIN:=$(GOPATH)/bin
LDFLAGS:="-X github.com/jorgechato/acictl/utils.VERSION=${MAIN_VERSION}"


.PHONY: release default test build run linter cov unit clean tools

default: test build

tools:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	go get -u github.com/jstemmer/go-junit-report

build:
	go build -ldflags ${LDFLAGS} -a -o acictl main.go

run:
	go run -ldflags ${LDFLAGS} main.go

clean:
	rm -rf acictl *.out *.xml release

release:
	./lazy/build.sh ${LDFLAGS}

test: fmtcheck
	go list $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=60s -parallel=4

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/lazy/gofmtcheck.sh'"

cov:
	go test -coverprofile=coverage.out ./...

unit:
	go test -v ./... | go-junit-report > test.xml

tools-dev: tools
	go get -u gopkg.in/alecthomas/gometalinter.v2
	gometalinter.v2 --install

linter:
	gometalinter.v2 --checkstyle > report.xml
