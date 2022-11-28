.PHONY: gen-go package

IMAGE_NAME=harbor.ctyuncdn.cn/ecf/protos
VERSION=v0.0.3
GIT_COMMIT=$(shell git rev-parse HEAD)

IMPORT_IMAGE_NAME=harbor.ctyuncdn.cn/ecf/goimports
IMPORT_VERSION=v0.0.1

default: gen-proto

gen-proto: clean
	docker run --rm -it -v `pwd`:/opt/protos ${IMAGE_NAME}:${VERSION} sh gen.sh proto
	docker run -v `pwd`:/code ${IMPORT_IMAGE_NAME}:${IMPORT_VERSION} -format-only -l -w /code/goout

clean:
	@rm -rf goout
