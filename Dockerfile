FROM quay.io/doliveira1277/golang-demo-base:v1.0 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o hello-world-api

FROM quay.io/doliveira1277/golang-demo-runtime:v1.0

WORKDIR /app/

COPY --from=builder /app/hello-world-api .

EXPOSE 8080

CMD ["./hello-world-api"]
