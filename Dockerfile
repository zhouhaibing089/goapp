FROM golang:1.20
WORKDIR /go/src/github.com/zhouhaibing089/goapp
ADD . .
RUN go install github.com/zhouhaibing089/goapp

FROM ubuntu:22.04
COPY --from=0 /go/bin/goapp /usr/local/bin/goapp
