# Step 1: Build Stage
FROM golang:1.24 AS builder

WORKDIR /app

# Copy go.mod and go.sum first (for dependency caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the project files
COPY . .

# Ensure a statically linked binary for Alpine Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o delivery-optimiser cmd/main.go

# Debug: Check if binary exists
RUN ls -lh /app

# Step 2: Runtime Stage
FROM alpine:latest

WORKDIR /app

# Install bash (optional for debugging)
RUN apk add --no-cache bash

# Copy compiled binary
COPY --from=builder /app/delivery-optimiser .

# Ensure binary has execution permissions
RUN chmod +x delivery-optimiser

# Create output directory
RUN mkdir -p /app/output

# Use volume for persistent output
VOLUME [ "/app/output" ]

# Expose output directory for mapped storage
CMD ["./delivery-optimiser"]