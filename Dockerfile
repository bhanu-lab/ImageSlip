FROM golang:1.17.1-alpine3.14 as builder

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app/ImageSlip/

RUN pwd

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN cd server && go build -o gRPCServer

RUN cd webserver && go build -o webserver

FROM alpine:3.14 as baseImage

# Installing required tools
RUN apk --update add supervisor

WORKDIR /app/ImageSlip/

COPY --from=builder /app/ImageSlip/server/gRPCServer .
COPY --from=builder /app/ImageSlip/webserver/webserver .
COPY --from=builder /app/ImageSlip/supervisor/service_script.conf supervisor/

# This container exposes port 8080 to the outside world
EXPOSE 8080
EXPOSE 39298

# Run the binary program produced by `go build`
CMD ["supervisord","-c","supervisor/service_script.conf"]
