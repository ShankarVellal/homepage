# Homepage
Get Go dependency:
- ``go get "golang.org/x/crypto/acme"


Build a Go static binary:
- ``cd go/src/main``
- ``CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .``

Run the static binary in a Docker container:
- ``docker build -t homepage-v1.2 .``
- ``docker run --name homepage -d -p 443:443 -t homepage-v1.2``