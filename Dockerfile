# Use the official Golang image
FROM golang:1.21

# Set the Working Directory
WORKDIR /go/src/app

# Copy the local code to the container
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Command to run the executable
CMD ["go", "run", "."]
