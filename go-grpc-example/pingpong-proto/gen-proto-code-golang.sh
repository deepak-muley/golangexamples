#!/usr/bin/env bash

PROTOC=`which protoc`

mkdir -p $(pwd)/generated/grpc

#compile .proto
protoc -I=$(pwd) --go_out=plugins=grpc:$(pwd)/generated/grpc $(pwd)/pingpong.proto
