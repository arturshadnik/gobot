# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.13 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY .  ./

WORKDIR /app/cmd/webserver
# Build the binary.
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=readonly -v -o gcp-data-drive

# Use the Google Distroless image for a minimal container.
FROM gcr.io/distroless/static

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/cmd/webserver/gcp-data-drive /gcp-data-drive

# Run the web service on container startup.
CMD ["/gcp-data-drive"]