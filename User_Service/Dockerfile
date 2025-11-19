# Use official Go image
FROM golang:1.25

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build the Go app
RUN go build -o user_service .

# Expose backend port
EXPOSE 8080

# Run the executable
CMD ["./user_service"]
