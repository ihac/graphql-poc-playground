FROM golang:1.16.5-alpine AS builder
COPY . /app
WORKDIR /app/reviews
RUN go build -o reviews

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/reviews .
ENTRYPOINT [ "/app/reviews" ]