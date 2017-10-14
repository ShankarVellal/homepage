FROM golang

ADD ./go/ .
COPY ./public/ ./public/
RUN go get golang.org/x/crypto/acme
RUN go install main
ENTRYPOINT /go/bin/main

EXPOSE 443 80