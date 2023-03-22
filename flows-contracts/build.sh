#!/bin/bash

for cmd in "protoc" "protoc-gen-go" "protoc-gen-go-grpc"; do
  type $cmd >/dev/null 2>&1 || { echo >&2 "$cmd required but it's not installed; aborting."; exit 1; }
done

PROTO_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
echo $PROTO_DIR

echo "compiling protofiles listed: "
for dir in $(find ${PROTO_DIR} -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq); do
  file=$(find "${dir}" -name '*.proto')
  echo "${file}"
  files="${files}${file} "
done
protoc -I ${PROTO_DIR} --go_out=paths=source_relative:${PROTO_DIR}/ \
--go-grpc_out=paths=source_relative:${PROTO_DIR}/ ${files} --experimental_allow_proto3_optional
