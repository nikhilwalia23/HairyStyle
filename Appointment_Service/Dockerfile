# Use a Go version that satisfies go.mod (1.23+)
FROM golang:1.23.3-alpine

# Set working directory
WORKDIR /app

# Copy dependency files first for caching
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go app
RUN go build -o appointment_service .

# Expose the port your app uses
EXPOSE 4242

# Run the app
CMD ["./appointment_service"]
