# Simple homepage using a Go web server

- ``docker build -t homepage-v1.2 .``
- ``docker run --name homepage -d -p 443:443 -p 80:80 -t homepage-v1.2``