#!/bin/bash

PROTOS=`find ./protos/pb -name '*.proto'`
PROTO_DIRS=$(echo ${PROTOS} | xargs -n1 dirname | sort -u)
OUT="generated/go/"

mkdir -p $OUT

IDL_PATH=/go/src/github.com/taeho-io/family/idl

for FILE in ${PROTOS}; do
  FILE=$IDL_PATH${FILE#.}

  # generate gRPC stub
  docker run --rm -it --name protoc -v $(pwd):$IDL_PATH -w $IDL_PATH xissy/protoc:v0.0.5 \
    -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I/go/src/github.com/grpc-ecosystem/grpc-gateway \
    -I$IDL_PATH/protos \
    --go_out=plugins=grpc:/go/src \
    $FILE

  # generate reverse-proxy
  docker run --rm -it --name protoc -v $(pwd):$IDL_PATH -w $IDL_PATH xissy/protoc:v0.0.5 \
    -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I/go/src/github.com/grpc-ecosystem/grpc-gateway \
    -I$IDL_PATH/protos \
    --grpc-gateway_out=logtostderr=true:/go/src \
    $FILE
done

# generate mocks
mockery -all -dir ./generated/go/pb/family -inpkg
