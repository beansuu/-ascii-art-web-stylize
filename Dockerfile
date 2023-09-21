# Start from a base image.
FROM golang:alpine AS build-env

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies.
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o ascii-art-web

# Start a new stage from scratch
FROM alpine

WORKDIR /app

# Copy the executable
COPY --from=build-env /app/ascii-art-web /app

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./ascii-art-web"]

LABEL version="1.0" \
      description="ASCII Art Web Server" \
      maintainer="beansu<mauno.talli@hotmail.com>"
