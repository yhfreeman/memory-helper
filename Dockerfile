# Start from the official Golang image
FROM golang:alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY ./src ./src
COPY ./static ./static

# Initialize the Go module
RUN go mod init memoryhelper # Replace 'memoryhelper' with your desired module name

# Get dependencies
RUN go get github.com/gin-gonic/gin
RUN go get github.com/garyburd/redigo/redis

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o memoryhelper ./src

# Start a new stage from scratch
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/memoryhelper .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./memoryhelper"]
