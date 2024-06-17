FROM golang:alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod tidy && go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -v -o main ./cmd/main.go

# Start a new stage from scratch
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy the executable file from the builder stage
COPY --from=builder /app/ .

ENV PORT 8080
EXPOSE $PORT

# Command to run the executable
ENTRYPOINT ["./main"]