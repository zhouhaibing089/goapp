FROM golang:1.20
WORKDIR /go/src/github.com/zhouhaibing089/k8s-kubelet-race
ADD . .
RUN go install github.com/zhouhaibing089/k8s-kubelet-race

FROM ubuntu:22.04
COPY --from=0 /go/bin/k8s-kubelet-race /usr/local/bin/k8s-kubelet-race