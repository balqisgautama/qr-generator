# Build stage
FROM golang:1.23-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s' -o qr-generator ./cmd/main.go

# Final stage
FROM alpine:latest

# Install necessary packages
RUN apk add curl

WORKDIR /home

# Copy the built binary and other necessary files
COPY --from=builder /app/qr-generator ./
COPY --from=builder /app /app

# Expose the application port
EXPOSE 8000

# Run the application with nodemon
CMD ["/home/qr-generator"]