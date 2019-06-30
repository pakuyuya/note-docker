FROM golang:1.11.11-alpine as builder
ENV CGO_ENABLED=0
COPY ./src /app
WORKDIR /app
RUN apk update \
    && apk add --no-cache git \
    && go build

FROM alpine:3.10
COPY --from=builder /app/sample-batch /app/sample-batch
COPY ./sql /app/sql
RUN apk add --no-cache ca-certificates

