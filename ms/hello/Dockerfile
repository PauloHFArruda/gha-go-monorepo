FROM golang:1.21.6-alpine AS builder

ENV MICROSERVICE_NAME=hello

WORKDIR /build

ADD ./ .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on
RUN go build -o bin/ms ms/hello/main.go

# RELEASE
FROM alpine:3.12

COPY --from=builder ["/build", "/"]

ENTRYPOINT ["./bin/ms"]