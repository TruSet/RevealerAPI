FROM golang:1.10
MAINTAINER Greg Taschuk
WORKDIR /go/src/github.com/TruSet/RevealerAPI

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . .

CMD "go run main.go"
