# goバージョン
FROM golang:1.19.1-alpine

RUN apk update && apk add git

# RUN go install github.com/go-delve/delve/cmd/dlv@latest

# RUN go install github.com/cosmtrek/air@latest
