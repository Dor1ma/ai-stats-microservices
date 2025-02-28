FROM golang:1.23.6-alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR /app

COPY proto ./proto
COPY service1 ./service1
COPY ../.env ./service1/.env

WORKDIR /app/service1
RUN go mod download
RUN go build -o ./bin/main ./cmd/service1/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/service1/bin/main .
COPY --from=builder /app/service1/.env .env

EXPOSE 50051
CMD ["./main"]