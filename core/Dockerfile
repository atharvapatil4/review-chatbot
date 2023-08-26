# Use a base image with Go pre-installed
FROM golang:latest

WORKDIR /app

COPY . /app
RUN go build
# Build and run the Go application
CMD ["./server"]