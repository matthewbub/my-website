# Use the latest Golang image as the base for the build-stage
FROM golang:1.20 AS build-stage

# Set the working directory inside the container to /app
WORKDIR /app

# Copy the Go module and sum files
COPY go.mod go.sum ./

# Download the Go modules specified in go.mod and go.sum
RUN go mod download

# Copy the entire project into the container
COPY . /app

# Compile the Go project to a binary named 'main'
RUN go build -o main .

# Specify the command to run the compiled binary
CMD ["./main"]
