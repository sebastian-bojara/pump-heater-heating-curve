# Build stage
FROM golang:1.25-alpine3.23 AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy source code
COPY main.go ./

# Build the application
RUN go build -o heating-curve .

# Final stage
FROM alpine:3.23

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/heating-curve .

# Set the entrypoint
ENTRYPOINT ["./heating-curve"]
