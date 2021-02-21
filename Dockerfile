FROM golang:1.15.8-alpine3.12

RUN apk add git

COPY . /go/src/github.com/flickyiyo/pokemon-api
WORKDIR /go/src/github.com/flickyiyo/pokemon-api

RUN go get
RUN go build


