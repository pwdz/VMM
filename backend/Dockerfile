# Use the official Go image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Build the Go application
RUN ls
RUN go build -o main .

# Expose the port the application will run on
EXPOSE 8000

# Run the Go application
CMD ["./main"]
