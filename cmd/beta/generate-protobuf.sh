#!/bin/bash

set -e

echo "Generating gRPC server"
protoc -I/usr/local/include \
  -I. \
  -I$GOPATH/src \
  --go_out=plugins=grpc:. \
  service/beta.proto
