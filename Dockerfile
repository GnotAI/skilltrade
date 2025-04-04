# Use the official Golang image as a base
FROM golang:1.23.5 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main ./main.go

# Start a new stage from scratch
FROM debian:bullseye-slim

# Install necessary dependencies for the application (if any)
RUN apt-get update && apt-get install -y ca-certificates && apt-get clean

# Copy the Pre-built binary file from the builder stage
COPY --from=builder /app/main /main

# Expose port 8080 (change this if your app uses a different port)
EXPOSE 3000

# Command to run the executable
CMD ["/main"]
