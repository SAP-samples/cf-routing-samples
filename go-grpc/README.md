# gRPC APP

App that serves gPRC traffic. Note that this will only work if all network hops
between the client and the app use HTTP/2. Configure your load balancers
appropriately.

## Running Locally
1. `go build`
2. `PORT=8080 ./grpc`
3. `grpcurl -vv -plaintext -import-path ./test -proto test.proto localhost:8080 test.Test.Run`

### Renerating code

```shell
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
protoc --go_out=$PWD example.proto
protoc --go-grpc_out=$PWD example.proto
```
