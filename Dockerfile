# Stage 1: Build the Go application
FROM golang:1.22.5-alpine as build

# Install CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon@latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o crud-in-go .

# Stage 2: Create the runtime image


# Stage 2: Create the runtime image
FROM alpine:latest

# Install necessary packages
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=build /app/crud-in-go .
COPY --from=build /go/bin/CompileDaemon /usr/local/bin/CompileDaemon

# Verify the presence of CompileDaemon and path
RUN echo $PATH

RUN ls -l /usr/local/bin/CompileDaemon

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["CompileDaemon", "-command=./crud-in-go"]




