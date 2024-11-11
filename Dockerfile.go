# Dockerfile.go

# Start with a lightweight Go image
FROM docker.io/golang:1.23.1-alpine

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go binary
RUN go build -o api main.go

# Expose the API port (adjust if different)
EXPOSE 8020

# Run the Go API
CMD ["./api"]
