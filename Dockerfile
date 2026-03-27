# Stage 1: Build stage
FROM golang:1.23-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build the application with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s -X main.version=$(git describe --tags 2>/dev/null || echo 'dev')" \
    -o /app/bin/api \
    ./cmd

# Stage 2: Final stage
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata curl

# Create non-root user
RUN addgroup -g 1000 -S appgroup && \
    adduser -u 1000 -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/bin/api /app/api

# Copy configuration files (if any)
# COPY --chown=appuser:appgroup configs/ /app/configs/

# Copy templates/assets (if any)
# COPY --chown=appuser:appgroup templates/ /app/templates/
# COPY --chown=appuser:appgroup static/ /app/static/

# Set permissions
RUN chmod +x /app/api && \
    chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8080

# Health check (optional - remove if using compose healthcheck)
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8080/health || exit 1

# Run the application
# CMD ["/app/api"]