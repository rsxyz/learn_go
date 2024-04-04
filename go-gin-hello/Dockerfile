# Start from the official Go image
FROM golang:1.21-alpine as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules and build files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o hello ./cmd/hello/main.go

# Start a new stage from scratch
FROM alpine:latest  

# Set the current working directory inside the container
WORKDIR /app

# Copy the built executable from the previous stage
COPY --from=builder /app/hello .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./hello"]
