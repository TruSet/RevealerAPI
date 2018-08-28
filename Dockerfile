FROM golang:1.9
MAINTAINER Greg Taschuk
WORKDIR /go/src/github.com/TruSet/RevealerAPI
RUN go get -u github.com/kardianos/govendor
# hacky way to get libsec bindings
# https://github.com/ethereum/go-ethereum/issues/2738
COPY vendor vendor
RUN go get -u github.com/ethereum/go-ethereum
#COPY "${GOPATH}/src/github.com/ethereum/go-ethereum/crypto/secp256k1/libsecp256k1" "vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/"
RUN govendor install +vendor,^program
RUN govendor sync
COPY . .
CMD "go run main.go -e docker"
