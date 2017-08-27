FROM golang

ADD ./go/ .
RUN go get golang.org/x/crypto/acme
RUN go install main
ENTRYPOINT /go/bin/main

EXPOSE 443