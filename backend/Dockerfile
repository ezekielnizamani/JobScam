# Use the official Golang image
FROM golang:1.23-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port on which the app will run
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
