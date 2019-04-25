FROM golang
ADD . /go/src/github.com/innovationnorway/go-hello-world
RUN go install github.com/innovationnorway/go-hello-world
ENTRYPOINT /go/bin/go-hello-world
