# Use the official Golang image
FROM golang:1.23

# Set the working directory
WORKDIR /app

# Copy the source code
COPY . .

# Install dependencies
RUN go mod tidy

# Build the application
RUN go build -o product-service main.go

# Expose the service port
EXPOSE 8081

# Run the application
CMD ["./product-service"]