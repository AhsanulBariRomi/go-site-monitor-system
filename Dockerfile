# Start from a small, secure base image
FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code (go.mod and main.go)
COPY . .

# Build the application
# This creates a binary executable named "monitor"
RUN go build -o monitor main.go

# Command to run the executable when the container starts
CMD ["./monitor"]