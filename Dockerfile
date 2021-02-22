FROM golang:1.15.8-alpine3.12 as build

RUN apk add git

COPY . /go/src/github.com/flickyiyo/pokemon-api
WORKDIR /go/src/github.com/flickyiyo/pokemon-api

RUN go get
RUN go build -o pokemon_api .

FROM alpine:latest as prod
WORKDIR /root
COPY --from=build /go/src/github.com/flickyiyo/pokemon-api/pokemon_api .
CMD [ "./pokemon_api" ]


