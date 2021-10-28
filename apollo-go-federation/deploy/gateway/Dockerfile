FROM golang:1.16.5-alpine AS builder
COPY . /app
WORKDIR /app/gateway
RUN go build -o gateway

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/gateway .
ENTRYPOINT [ "/app/gateway" ]