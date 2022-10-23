
######################################
#   STEP 1 build executable binary   #
######################################
FROM golang:1.19.2-alpine3.16 AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR /go/src/github.com/havus/basic-webhook-server
COPY . .

# Fetch dependencies:
# - Using go get.
RUN go get -d -v

# Build the binary.
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/basic-webhook-server

##################################
#   STEP 2 build a small image   #
##################################
FROM alpine:3.16.2
WORKDIR /go/bin/

# Copy our static executable.
COPY --from=builder /go/bin/basic-webhook-server .
COPY .env .

# Run the basic-webhook-server binary.
CMD ["./basic-webhook-server"]
