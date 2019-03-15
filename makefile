.DEFAULT_GOAL := build

test:
	go test -v -race ./...

build:
	go build

all: test build
