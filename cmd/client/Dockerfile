# Start from the official Golang base image
FROM golang:1.23 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files (i am on path git-test/cmd/client, but the go.mod and go.sum files are in the root of the project)
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o client ./cmd/client

# Start a new stage from scratch
FROM debian:bookworm-slim

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/client /app/client

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/app/client"]