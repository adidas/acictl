#!/usr/bin/env bash

set -e
echo "" > coverage.xml

for D in $(go list ./... | grep -v vendor | grep -v cmd); do
    go test -coverprofile=coverage.out $D
done
