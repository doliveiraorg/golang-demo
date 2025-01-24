FROM golang:1.24rc2-alpine3.21 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o hello-world-api

FROM quay.io/fedora/httpd-24-micro:latest

WORKDIR /app/

COPY --from=builder /app/hello-world-api .

EXPOSE 8080

CMD ["./hello-world-api"]
