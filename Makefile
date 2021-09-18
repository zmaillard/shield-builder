PROJECT_NAME ?= roadsign-shield-builder

clean:
	rm -rf ./bin ./vendor Gopkg.lock

buildfn:
	export GO111MODULE=on
	 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./azure-function/main main.go rice-box.go

build:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o ./azure-function/main main.go rice-box.go


.PHONY: build clean
