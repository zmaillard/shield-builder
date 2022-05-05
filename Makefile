PROJECT_NAME ?= roadsign-shield-builder
ENV ?= dev

build:
	go build -o ./bin/shield main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

default: build

.PHONY: build clean default
