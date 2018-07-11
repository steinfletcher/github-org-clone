.DEFAULT_GOAL := build

test:
	go test -v -race ./...

build:
	go build

install:
	go install

all: test build install
