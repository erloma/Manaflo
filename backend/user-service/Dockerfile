# Stage 1: Build the Go binary
FROM golang:1.24-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules files
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN GOARCH=amd64 GOOS=linux go build -o main .

# Stage 2: Create a small image with only the binary
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy the environment variables
COPY .env /root/.env

# Expose port 8080
EXPOSE 8080

# Command to run the binary
CMD ["./main"]
