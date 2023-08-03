# Go gRPC Sample

Example gRPC server implemented in Go
## Deploying the Go app
### Deploy the app using the manifest file
Deploy the app
```shell
> export DOMAIN=my.cf.app.domain
> cf push -f app-manifest.yml --var domain=$DOMAIN
```

### Deploy the app without the manifest file
```shell
> cf push --no-route go-grpc-test --buildpack go_buildpack 
# my.cf.app.domain is used as an example for demonstration purpose
> cf map-route go-grpc-test my.cf.app.domain --hostname go-grpc-test --app-protocol http2
```

## Testing the Go app
`grpcurl` needs to be [installed separately](https://github.com/fullstorydev/grpcurl).
```shell
> grpcurl go-grpc-test.my.cf.app.domain:443 example.Example.Run 
```

## Building and running on the local machine

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
