FROM quay.io/doliveira1277/golang-demo-base:v1.0 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o demo-app

FROM quay.io/doliveira1277/golang-demo-runtime:v1.0

WORKDIR /app/

COPY --from=builder /app/demo-app .

COPY static/ /app/static/

EXPOSE 8080

CMD ["./demo-app"]
