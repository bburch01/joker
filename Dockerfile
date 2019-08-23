# joker/Dockerfile

FROM golang:1.12
FROM golang:alpine as builder

RUN apk --no-cache add --update git gcc libc-dev net-tools

WORKDIR /app/go/src/github.com/bburch01/joker
COPY . .

ENV GOPATH /app/go
ENV GO111MODULE=auto

RUN go get github.com/gorilla/mux
RUN go mod download
RUN GOOS=linux go build -i -a -o joker

FROM alpine:latest

RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache curl

COPY --from=builder /app/go/src/github.com/bburch01/joker/joker .

ENTRYPOINT ./joker --port 8080
