# Build go app and start the http server of the calculator on port 80
# Use the official Golang base image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the Go server files to the container
COPY main.go go.mod go.sum /app/
COPY pkg /app/pkg/

# Build the Go server executable
RUN go build -o server

# Expose the port that the server listens on
EXPOSE 80

# Run the Go server when the container starts
CMD ["./server"]
