# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Moez Gharsallah <gharsallahmoez@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux make build

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk update && apk add --no-cache ca-certificates bash git

WORKDIR /app/

COPY config/config.prod.yml config/config.dev.yml ./config/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/bin/chat-cli-go .
RUN  chmod +x chat-cli-go
EXPOSE 3000
# Command to run the executable
CMD ["./chat-cli-go"]