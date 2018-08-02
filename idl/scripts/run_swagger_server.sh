#!/usr/bin/env bash

echo "Server will be running on http://localhost:80..."

docker run --rm -it \
    --name=taeho-io-idl-swagger-server \
    -p 80:8080 \
    -e SWAGGER_JSON=/swagger/swagger.json \
    -v $GOPATH/src/github.com/taeho-io/family/idl/generated/swagger:/swagger \
    swaggerapi/swagger-ui:3.16.0
