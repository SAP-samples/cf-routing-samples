# gRPC Example App

Example gRPC server implemented in Ruby

## Running Locally
1. `bundle install`
2. `PORT=8080 bundle exec ruby server.rb`

### Renerating code

```shell
bundle install
bundle exec grpc_tools_ruby_protoc -I . --ruby_out . --grpc_out . example.proto
```
