# Ruby gRPC Sample

Example gRPC server implemented in Ruby

## Running Locally
1. `bundle install --path vendor/bundle`
2. `PORT=8080 bundle exec ruby server.rb`
3. `grpcurl -proto example.proto -plaintext localhost:8080 Example.Run`

### Renerating code

```shell
bundle install --path vendor/bundle
bundle exec grpc_tools_ruby_protoc -I . --ruby_out . --grpc_out . example.proto
```
