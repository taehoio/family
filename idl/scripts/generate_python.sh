#!/bin/bash

PROTOS=`find ./protos/pb -name '*.proto'`
PROTO_DIRS=$(echo ${PROTOS} | xargs -n1 dirname | sort -u)
OUT="generated/python/"

mkdir -p $OUT

IDL_PATH=/go/src/github.com/taeho-io/family/idl

for FILE in ${PROTOS}; do
  FILE=$IDL_PATH${FILE#.}

  # generate gRPC stub
  docker run --rm -it --name grpcio-tools -v $(pwd):$IDL_PATH -w $IDL_PATH xissy/grpcio-tools:v0.0.1 \
    -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I/go/src/github.com/grpc-ecosystem/grpc-gateway \
    -I$IDL_PATH/protos \
    --python_out=:$OUT \
    --grpc_python_out=$OUT \
    $FILE
done
