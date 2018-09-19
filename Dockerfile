FROM golang:1.9
MAINTAINER Greg Taschuk
WORKDIR /go/src/github.com/TruSet/RevealerAPI

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

#RUN go get -u github.com/kardianos/govendor
#COPY vendor vendor
# hacky way to get libsec bindings
# https://github.com/ethereum/go-ethereum/issues/2738
#RUN go get -u github.com/ethereum/go-ethereum
#RUN dep init
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . .
#RUN dep ensure
#--vendor-only

#RUN govendor add +external
#RUN cp -r ${GOPATH}/src/github.com/ethereum/go-ethereum/crypto/secp256k1/libsecp256k1 ./vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/.
#RUN govendor sync
CMD "go run main.go -e docker"
