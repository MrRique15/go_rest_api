#Dockerfile
FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy
RUN go mod download

COPY . .

RUN go build -o /app/inventory_service . 

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/inventory_service .

CMD ["./inventory_service"]