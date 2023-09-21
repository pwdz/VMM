# Use an official Golang runtime as a parent image
FROM golang:1.20

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go application
RUN go build -o myapp

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]