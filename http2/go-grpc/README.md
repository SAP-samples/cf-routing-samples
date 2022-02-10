# Go gRPC Sample

Example gRPC server implemented in Go

## Running Locally

1. `go build`
2. `PORT=8080 ./go-grpc`
3. `grpcurl -vv -plaintext localhost:8080 example.Example.Run`

`grpcurl` needs to be [installed separately](https://github.com/fullstorydev/grpcurl).

### Renerating code

The generated code is created using the Protoc compiler, which needs to be [installed separately](https://grpc.io/docs/protoc-installation/).

The following steps can then be used to re-generate the gRPC files in `example/`.

```shell
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
protoc --go_out=$PWD example.proto
protoc --go-grpc_out=$PWD example.proto
```
