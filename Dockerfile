# Start from golang v1.16 base image
FROM golang:1.16.3-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/TDBackend

# Install the librdkafka
RUN apk -U add ca-certificates
RUN apk update && apk upgrade && apk add --no-cache curl pkgconfig bash build-base

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Build the executable
RUN GOARCH=amd64 go build -tags musl -a -installsuffix cgo -o TDBackend .

### Creating final image
# Creating a smaller image
FROM alpine:latest

# Add Maintainer Info
LABEL maintainer="TriDubai"

# ENV variable for TriDubai Directory
ENV APPDIR="/TriDubai"

# Create TriDubai folders
RUN mkdir -p $APPDIR/config $APPDIR/static

# Copy data from builder to target docker image
COPY --from=builder /go/TDBackend/config/*.yml $APPDIR/config/
COPY --from=builder /go/TDBackend/static/* $APPDIR/static/
COPY --from=builder /go/TDBackend/TDBackend $APPDIR/TDBackend

# Set the Current Working Directory inside the container
WORKDIR $APPDIR

# This container exposes port 8080 for rest APIs
EXPOSE 8080

# Run the executable
CMD ["./TDBackend"]