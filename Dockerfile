# Use specific versions for both build stages to ensure consistency and reproducibility
FROM node:20.11.1 AS node-build-stage

WORKDIR /app

# Copying package.json and package-lock.json separately allows Docker to cache installed dependencies unless these files change
COPY apps/com/package*.json ./

RUN npm install

# Copy the rest of the application code
COPY apps/com/ .

# Build the application
RUN npm run build

# Start a new build stage for the Go application
FROM golang:1.20 AS go-build-stage

WORKDIR /app

# Copying go.mod and go.sum separately allows Docker to cache downloaded modules unless these files change
COPY go.mod go.sum ./

RUN go mod download

# Copy the Go source code and the built assets from the previous stage
COPY . .
COPY --from=node-build-stage /app/dist /app/dist

# Compile the Go application
RUN go build -o main .

# Final stage: Use a smaller base image if possible to reduce the final image size
FROM golang:1.20

WORKDIR /app

# Copy the compiled binary from the build stage
COPY --from=go-build-stage /app/main .

# Adjusted destination path to match the expected structure by the Go application
COPY --from=node-build-stage /app/dist ./apps/com/dist

# Specify the command to run the compiled binary
CMD ["./main"]