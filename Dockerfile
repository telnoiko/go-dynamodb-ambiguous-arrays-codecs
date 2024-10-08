# Dockerfile for local builds
FROM golang:1.22-alpine

WORKDIR /src

COPY . .

RUN go build -o /test-app .

ENV ENV=local

ENTRYPOINT ["/test-app"]
