# Use the official Golang image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the local source code into the container's working directory
COPY . .

# Build the Golang binary executable
RUN go build -o central-api cmd/central-api/main.go

# Expose the port your API is running on (change 8080 to your API's actual port)
EXPOSE 8080

# Command to run the API when the container starts
CMD ["./central-api"]