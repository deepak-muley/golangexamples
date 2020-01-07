#!/bin/bash -x

pushd ./pingpong-proto
./gen-proto-code-golang.sh
popd

#TODO Move to Gopkg.toml
go get -u google.golang.org/grpc

go build ./...
go vet ./...
go install ./...