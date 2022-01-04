# gRPC APP

Example gRPC server implemented in Go

## Running Locally
1. `go build`
2. `PORT=8080 ./grpc`
3. `grpcurl -vv -plaintext localhost:8080 example.Example.Run`

### Renerating code

```shell
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
protoc --go_out=$PWD example.proto
protoc --go-grpc_out=$PWD example.proto
```
