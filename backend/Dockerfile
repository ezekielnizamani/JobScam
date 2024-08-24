# Use the official Golang image
FROM golang:1.23-alpine

# Install CompileDaemon for hot reloading
RUN go install github.com/githubnemo/CompileDaemon@latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Install any system dependencies if needed
RUN apk add --no-cache bash

# Install go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code (initial setup, not necessary with bind mount)
COPY . .

# Expose the port the app runs on
EXPOSE 8080

# Run the application with hot reloading
CMD ["CompileDaemon", "--build=go build -o main .", "--command=./main"]
