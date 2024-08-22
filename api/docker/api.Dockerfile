# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.22.5

ARG VERSION
ARG MODULE_NAME

WORKDIR /app

# Copy local code to the container image.
COPY . ./

WORKDIR /app/src
RUN go mod download

# Build the binary.
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -v -o bin/api \
    -ldflags "-X main.version=${VERSION}" \
    ${MODULE_NAME}/cmd

# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=0 /app/src/bin/api /bin/api

# Run the web service on container startup.
CMD ["/bin/api"]
