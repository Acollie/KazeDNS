# Use the golang image with the --platform flag to specify the target platform
FROM --platform=$BUILDPLATFORM golang:latest

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Set the target architecture for the Go build
ARG TARGETPLATFORM
RUN if [ "$TARGETPLATFORM" = "linux/arm64" ]; then \
      GOARCH=arm64; \
    elif [ "$TARGETPLATFORM" = "linux/arm/v7" ]; then \
      GOARCH=arm; \
    else \
      GOARCH=amd64; \
    fi && \
    go build -o dns -ldflags="-s -w" main.go

# Expose ports dns server
EXPOSE 8888 8888

# Used for API
EXPOSE 2112 2112

# Run the application
CMD ["./dns"]