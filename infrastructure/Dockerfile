# Use the official Golang image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod ./
COPY go.sum ./

# Download and install Go dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go API executable
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
