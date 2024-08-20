#Dockerfile
FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/main_api . 

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main_api .

CMD ["./main_api"]