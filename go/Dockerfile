# Stage 1: Build
FROM golang:1.20 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum for dependency installation
COPY go.mod go.sum ./

# Download and cache the dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o prometheus-config-updater

# Stage 2: Run
FROM alpine:latest

# Install required packages (optional)
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/prometheus-config-updater .

# Expose a port if required (optional)
# EXPOSE 8080

# Command to run the binary
CMD ["./prometheus-config-updater"]
