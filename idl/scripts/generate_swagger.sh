#!/bin/bash

OUT="generated/swagger/"
mkdir -p $OUT

IDL_PATH=/go/src/github.com/taeho-io/family/idl

# generate Swagger JSON files for each gRPC service
PROTOS=`find ./protos/pb -name '*.proto'`
PROTO_DIRS=$(echo ${PROTOS} | xargs -n1 dirname | sort -u)

for FILE in ${PROTOS}; do
  FILE=$IDL_PATH${FILE#.}

  docker run --rm -it --name protoc -v $(pwd):$IDL_PATH -w $IDL_PATH xissy/protoc:v0.0.3 \
    -I/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I$IDL_PATH/protos \
    --swagger_out=logtostderr=true:$OUT \
    $FILE
done

# merge the Swagger JSON files to a Swagger JSON file
SWAGGERS=`find ./generated/swagger -name '*.swagger.json' | sort -u`
INPUT_FILES_OPTIONS=''
for FILE in ${SWAGGERS}; do
  FILE=${FILE#./generated}

  INPUT_FILES_OPTIONS=$INPUT_FILES_OPTIONS'--input '$FILE' '
done
docker run --rm -it \
    --name=taeho-io-idl-swagger-tool-merge \
    -v $GOPATH/src/github.com/taeho-io/family/idl/generated/swagger:/swagger \
    xissy/swagger-tool:v0.0.2 \
    merge $INPUT_FILES_OPTIONS -o /swagger/swagger.json
