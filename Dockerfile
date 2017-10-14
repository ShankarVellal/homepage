FROM golang

ADD ./go/ .
RUN go get golang.org/x/crypto/acme
RUN go install main
RUN cp /go/src/main/public /go/bin/ -R -a
ENTRYPOINT /go/bin/main

EXPOSE 443 80