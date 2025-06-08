# Use the official Golang image
FROM golang:1.23

# Set the working directory
WORKDIR /app

# Copy the source code
COPY . .

# Install dependencies
RUN go mod tidy

# Build the application
RUN go build -o user-service main.go

# Expose the service port
EXPOSE 8080

# Run the application
CMD ["./user-service"]