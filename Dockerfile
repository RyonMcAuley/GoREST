# 1. Build stage
FROM golang:1.24.3-alpine AS builder

WORKDIR /app

# Copy go files
COPY go.mod go.sum ./
RUN go mod download

# Copy everything else
COPY . .

# Build the Go app
RUN go build -o server ./cmd/server

# 2. Runtime stage
FROM alpine:latest

WORKDIR /root/

# Copy built binary and web files
COPY --from=builder /app/server .
COPY --from=builder /app/web ./web

# Port your server listens on
EXPOSE 8080

# Run the server
CMD ["./server"]
