FROM golang:alpine as builder

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
FROM alpine:latest


# Add tzdata pkg to support TZ env var
RUN apk add tzdata

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /app/server

# Declare container exposed port
EXPOSE 8080

# Run the web service on container startup.
CMD ["/app/server"]

