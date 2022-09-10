FROM golang:buster as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.COPY go.* .
COPY go.* ./
RUN go mod download

# Copy all code
COPY . ./

# Build binary
RUN go build -v -o server

#Stage 2
#######################################
# Use the official Debian slim image for a lean production container.
FROM debian:buster-slim

# Install ca-certifactes
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /app/server

# Declare container exposed port
EXPOSE 8080

# Run the web service on container startup.
CMD ["/app/server"]

