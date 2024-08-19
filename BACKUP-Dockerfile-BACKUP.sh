# Dockerfile
FROM golang:1.20-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY .env /app/.env

RUN go mod download

COPY . .

RUN go build -o /bin/main_api ./main_api

RUN go build -o /bin/saga_execution_controller ./saga_execution_controller

RUN go build -o /bin/inventory_service ./inventory_service

RUN go build -o /bin/payment_service ./payment_service

RUN go build -o /bin/shipping_service ./shipping_service

FROM alpine:latest

EXPOSE 8080

WORKDIR /root/

COPY --from=builder /bin /bin

CMD ["sh"]