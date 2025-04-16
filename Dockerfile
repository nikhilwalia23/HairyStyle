# Use a minimal base image
FROM golang:1.23.3 as builder

# Set working directory
WORKDIR /app

# Copy Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the Go binary
RUN go build -o server main.go

# Use a smaller base image for final container
FROM debian:bookworm-slim

# Set working directory
WORKDIR /app

# Copy compiled Go binary
COPY --from=builder /app/server .

# Copy TLS certificates (if using self-signed certs)
COPY cert.pem key.pem ./

# Expose HTTP/3 (UDP 8000)
EXPOSE 8800/udp
EXPOSE 8800/tcp

# Run the server
CMD ["./server"]
