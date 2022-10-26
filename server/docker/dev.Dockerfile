FROM golang:1.19

RUN mkdir /go/src/server
RUN mkdir /go/src/sqlite
WORKDIR /go/src/server

COPY go.mod go.sum ./
RUN go mod download