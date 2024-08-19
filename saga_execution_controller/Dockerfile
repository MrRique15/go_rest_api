#Dockerfile
FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy
RUN go mod download

COPY . .

RUN go build -o /app/saga_execution_controller . 

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/saga_execution_controller .

CMD ["./saga_execution_controller"]