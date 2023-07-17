GIT_VERSION := $(shell git describe --abbrev=0 --tags)

ifndef GIT_VERSION
GIT_VERSION = main
endif

default: build

test:
	go test ./...

build:
	go build -ldflags "-s -w -X main.version=${GIT_VERSION}"

install: build
	mkdir -p ~/.tflint.d/plugins
	mv ./tflint-ruleset-template ~/.tflint.d/plugins
