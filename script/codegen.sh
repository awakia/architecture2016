#!/bin/bash

mkdir -p src/rpc

# Generate service by gRPC
protoc -Iproto -I$GOPATH/src/github.com/gengo/grpc-gateway/third_party/googleapis --go_out=Mgoogle/api/annotations.proto=github.com/gengo/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:src/rpc proto/*.proto

# Generate reverse-proxy by grpc-gateway

protoc -Iproto -I$GOPATH/src/github.com/gengo/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:src/rpc proto/*.proto
