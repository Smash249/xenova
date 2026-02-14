# Build stage
FROM golang:1.25.6-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod ./

# Copy source code to discover dependencies
COPY . .

# Download dependencies and generate go.sum
RUN go mod download && go mod tidy

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

# Copy configuration files
COPY --from=builder /app/config ./config

# Expose port
EXPOSE 8099

# Run the application
CMD ["./main"]
