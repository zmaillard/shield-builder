PROJECT_NAME ?= roadsign-shield-builder
ENV ?= dev

build:
	go build -ldflags="-s -w"  -o ./bin/shield main.go rice-box.go
clean:
	rm -rf ./bin ./vendor Gopkg.lock

default: build

.PHONY: build clean default
