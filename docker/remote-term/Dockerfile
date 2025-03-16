FROM golang:1.23.4-bullseye

ADD . /go/src/github.com/josegonzalez/go-remote-term
WORKDIR /go/src/github.com/josegonzalez/go-remote-term

ENV GOOS=linux
ENV GOARCH=arm64
ENV CGO_ENABLED=1
ENV CC=aarch64-linux-gnu-gcc

RUN go build -tags nogcs,nos3 \
    -ldflags "-linkmode external -extldflags '--static-pie' -s -w" \
    -buildmode=pie \
    -o remote-term

RUN apt-get update && apt-get install -y file

RUN ls -lah /go/src/github.com/josegonzalez/go-remote-term
RUN file /go/src/github.com/josegonzalez/go-remote-term/remote-term
