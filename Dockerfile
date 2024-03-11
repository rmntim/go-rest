FROM golang:alpine as builder

# Important: required for go-sqlite3
ENV CGO_ENABLED=1
RUN apk add --no-cache \
    gcc \
    # Required for Alpine
    musl-dev

WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN go build -v -o server /app/cmd/url-shortener

# Use alpine to run the binary
FROM alpine

WORKDIR /app

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /app/server
COPY --from=builder /app/config/prod.yaml /app/config.yaml

EXPOSE 8081

# Run the web service on container startup.
CMD ["/app/server"]
