# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.13-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -ldflags="-s -w"  -o /main main.go rice-box.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /main /main

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/main"]