# Use the official Golang image as the base image
FROM golang:1.22.4 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /service-catalog

# Use a minimal base image
FROM alpine:3.17

# Set the working directory inside the container
WORKDIR /app

# Copy the built application from the builder stage
COPY --from=builder /service-catalog /app/service-catalog

# Expose the port the service will run on
EXPOSE 8080

# Command to run the application
CMD ["/app/service-catalog"]
