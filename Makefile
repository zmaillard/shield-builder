PROJECT_NAME ?= roadsign-shield-builder
ENV ?= dev


AWS_BUCKET_NAME ?= $(PROJECT_NAME)-artifacts-$(ENV)
AWS_STACK_NAME ?= $(PROJECT_NAME)-stack-$(ENV)
AWS_REGION ?= us-west-2

FILE_TEMPLATE = cloudformation.yml
FILE_PACKAGE = package.yml

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/shields viewmain.go rice-box.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/shieldsave savemain.go rice-box.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose -s dev

deployprod: clean build
	sls deploy --verbose -s prod

.PHONY: build clean deploy gomodgen
