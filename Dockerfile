FROM golang

ADD ./go/ .
ADD ./public/ .
RUN go get golang.org/x/crypto/acme
RUN go install main
ENTRYPOINT /go/bin/main

EXPOSE 443 80