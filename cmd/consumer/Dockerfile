# Build stage
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o consumer ./cmd/consumer

# Run stage
FROM debian:bookworm-slim

# Set the Current Working Directory inside the container
WORKDIR /app

# Install curl
RUN apt-get update && apt-get install -y curl && rm -rf /var/lib/apt/lists/*

# Copy the binary from the build stage.
COPY --from=builder /app/consumer /app/consumer

# Command to run the executable
CMD ["./consumer"]
