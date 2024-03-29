FROM golang:1.21 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY env.yaml .

EXPOSE 8080

CMD ["./main"]
