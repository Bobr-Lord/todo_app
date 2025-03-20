FROM golang:1.24-alpine3.21 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o server_app ./cmd/todo/main.go

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app .

EXPOSE 8080
EXPOSE 5432

CMD ["./server_app"]

