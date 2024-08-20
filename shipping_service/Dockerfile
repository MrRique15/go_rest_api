#Dockerfile
FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/shipping_service . 

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/shipping_service .

CMD ["./shipping_service"]