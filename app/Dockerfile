FROM golang:1.18-alpine

RUN apk update && apk add --no-cache git gcc musl-dev

RUN mkdir /go/src/app

WORKDIR /go/src/app

ADD . /go/src/app/

RUN go build *.go