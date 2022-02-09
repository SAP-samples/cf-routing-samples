# HTTP/2 H2C Sample

Example HTTP/2 server implemented in Go

The app supports HTTP/1.1 and HTTP/2 H2C if requested explicitly, i.e. no connection upgrade.

## Running Locally

1. `go build`
2. `PORT=8080 ./go-http2`
3. `curl -v --http2-prior-knowledge http://localhost:8080`
