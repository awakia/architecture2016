#!/bin/bash

mkdir -p src
protoc -Iproto -I$GOPATH/src/github.com/gengo/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:src proto/*.proto
