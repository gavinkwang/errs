#!/bin/bash

gen_proto() {

  rm -rf /tmp/protos
  mkdir -p /tmp/protos


  for f in $(find ./src -name "*.proto") ; do
    echo "gen proto File -> $f";
    protoc -Isrc -I. -I/usr/local/include \
    --go_out=plugins=grpc:/tmp/protos \
    $f
  done ;

  cp -rf /tmp/protos/gitlab.ctyuncdn.cn/rcr/status/protos/goout .
}


if [ "all" == "$2" ]; then
  gen_all $1
else
  gen_$1
fi;
