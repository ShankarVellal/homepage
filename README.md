# Homepage

Build a Go static binary:
- ``cd go/src/main``
- ``CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .``

Run the static binary in a Docker container:
- ``docker build -t homepage-v1 .``
- ``docker run --name homepage -d -p 80:8080 -t homepage-v1``