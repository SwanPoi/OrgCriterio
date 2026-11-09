# --- Этап 1: Сборка приложения ---
FROM golang:1.27-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/app/main.go

# --- Этап 2: Финальный легкий образ ---
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/server .
COPY internal/config/config.yaml ./config/config.yaml

EXPOSE 8080

CMD ["./server"]