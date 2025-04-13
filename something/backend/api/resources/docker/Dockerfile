#FROM public.ecr.aws/docker/library/golang:1.22.2-alpine3.19
FROM golang:1.22.2-alpine3.19

RUN apk update && \
    apk upgrade && \
    apk add make git bash

RUN go install github.com/cespare/reflex@latest
ENV CGO_ENABLED=0

WORKDIR /go/src/app

EXPOSE 8080

CMD reflex  -s -r '\.go$$' go run src/app/http/main.go

