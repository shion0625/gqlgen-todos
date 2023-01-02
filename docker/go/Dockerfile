# goバージョン
FROM golang:1.19.1-alpine

RUN apk update && apk add git && apk add alpine-sdk

ENV APP_ROOT /app

RUN mkdir $APP_ROOT

WORKDIR $APP_ROOT

ENV CGO_ENABLED=1 \
  GOOS=linux \
  GOARCH=amd64 \
  GO111MODULE=on

COPY . $APP_ROOT

EXPOSE 8080
