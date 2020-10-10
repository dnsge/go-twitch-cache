# Build stage
FROM golang:1.15-alpine AS build

WORKDIR /go/src/github.com/dnsge/go-twitch-cache

ADD go.mod .
ADD go.sum .
RUN go mod download

ADD . .
RUN go build -o /go/bin/twitch-cache-server ./cmd/twitch-cache-server

# Final stage
FROM alpine

WORKDIR /app
COPY --from=build /go/bin/twitch-cache-server /app/

ENTRYPOINT ["./twitch-cache-server"]
