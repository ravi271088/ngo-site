# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copy the dependency files
COPY go.mod go.sum* ./
COPY vendor/ ./vendor/

# Copy ALL source code first to avoid "directory not found" errors
COPY . .

# Build the application using the vendor folder
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o server ./cmd/server/main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/server .
# Copy web assets
COPY --from=builder /app/web ./web

EXPOSE 8080

CMD ["./server"]
