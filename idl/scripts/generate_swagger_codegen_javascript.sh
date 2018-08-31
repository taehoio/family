#!/bin/bash

IDL_PATH=/go/src/github.com/taeho-io/family/idl

docker run --rm -it \
    --name swagger-codegen \
    -v $(pwd):$IDL_PATH \
    -w $IDL_PATH \
    swaggerapi/swagger-codegen-cli:v2.3.1 \
        generate \
            -i generated/swagger/swagger.json \
            -l javascript \
            -o generated/swagger/client/javascript/
