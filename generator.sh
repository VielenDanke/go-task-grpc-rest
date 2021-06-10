#!/bin/sh -e

INC=$(go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway)
INC_OPENAPI=$(go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway/v2)
ARGS="-I${INC}"
ARGS_OPENAPI="-I${INC_OPENAPI}"

protoc $ARGS/third_party/googleapis $ARGS_OPENAPI -Iproto --grpc-gateway_out=logtostderr=true:. --swagger_out=allow_merge=true,merge_file_name=api:. \
 --go_out=plugins=grpc:. ./proto/*.proto