# Stage 1: Build the Go application
FROM golang:1.24 AS builder

# Set the working directory
WORKDIR /go-bank

# Copy go.mod and go.sum for dependency management
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o go-bank

# Stage 2: Create the final image
FROM alpine:latest

# Install necessary libraries if needed
RUN apk add --no-cache libc6-compat

# Copy the binary from the builder stage
COPY --from=builder /go-bank/go-bank .

# Make the binary executable
RUN chmod +x go-bank

# Command to run the application
CMD ["./go-bank"]
