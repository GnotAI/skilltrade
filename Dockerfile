# Use the official Go image as the base image
FROM golang:1.24.1-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 3000

# Command to run the executable
CMD ["./main"]
