FROM golang:1.13
MAINTAINER Neil McLaren & Greg Taschuk
WORKDIR /truset-revealer

# Gather dependencies before building the code, so that we don't re-gather them for every code change
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Now build
RUN go build ./...
CMD "go run main.go"
