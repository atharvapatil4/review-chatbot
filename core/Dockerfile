FROM golang:latest
WORKDIR /app

# Copy only the go module files for deps resolution
COPY go.mod go.sum ./

# Download dependencies to cache them in a Docker layer
RUN go mod download

# Copy the rest of the codebase
COPY . /app

# Build the Go application 
# On code changes w/o new deps, this is the only step that runs
RUN go build

CMD ["./server"]
