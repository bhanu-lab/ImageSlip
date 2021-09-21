FROM golang:1.17.1-alpine3.14

RUN apk add --no-cache git

# Installing required tools
RUN apk --update add supervisor

# Set the Current Working Directory inside the container
WORKDIR /app/ImageSlip/

RUN pwd

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# This container exposes port 8080 to the outside world
EXPOSE 8080
EXPOSE 39298

# Run the binary program produced by `go build`
CMD ["supervisord","-c","supervisor/service_script.conf"]
