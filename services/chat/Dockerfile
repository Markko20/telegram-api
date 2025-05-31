FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o chat-service main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/chat-service .

EXPOSE 8080

CMD ["./chat-service"]
