FROM golang:1.16.5-alpine AS builder
COPY . /app
WORKDIR /app/accounts
RUN go build -o accounts

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/accounts .
ENTRYPOINT [ "/app/accounts" ]