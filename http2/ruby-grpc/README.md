# Ruby gRPC Sample

Example gRPC server implemented in Ruby
## Deploying the Ruby app
### Deploy the app using the manifest file
Deploy the app
```shell
export DOMAIN=my.cf.app.domain
cf push -f app-manifest.yml --var domain=$DOMAIN
```

### Deploy the app without the manifest file
```shell
cf push --no-route ruby-grpc-test --buildpack https://github.com/cloudfoundry/ruby-buildpack
# my.cf.app.domain is used as an example for demonstration purpose
cf map-route ruby-grpc-test my.cf.app.domain --hostname ruby-grpc-test --app-protocol http2
```

## Testing the Ruby app
`grpcurl` needs to be [installed separately](https://github.com/fullstorydev/grpcurl).
```shell
grpcurl -proto example.proto ruby-grpc-test.my.cf.app.domain:443 Example.Run
```

## Building and running on the local machine
1. `bundle install --path vendor/bundle`
2. `PORT=8080 bundle exec ruby server.rb`
3. `grpcurl -proto example.proto -plaintext localhost:8080 Example.Run`

### Renerating code

```shell
bundle install --path vendor/bundle
bundle exec grpc_tools_ruby_protoc -I . --ruby_out . --grpc_out . example.proto
```
