FROM golang:1.23.6-alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR /app

COPY proto ./proto
COPY service2 ./service2
COPY ../.env ./service2/.env

WORKDIR /app/service2
RUN go mod download
RUN go build -o ./bin/main ./cmd/service2/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/service2/bin/main .
COPY --from=builder /app/service2/.env .env

EXPOSE 8080
CMD ["./main"]