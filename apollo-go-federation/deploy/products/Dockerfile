FROM golang:1.16.5-alpine AS builder
COPY . /app
WORKDIR /app/products
RUN go build -o products

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/products .
ENTRYPOINT [ "/app/products" ]