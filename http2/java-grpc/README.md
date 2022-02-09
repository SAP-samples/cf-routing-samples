# gRPC Example App

Example gRPC server implemented in Java

## Running Locally
1. `gradle build`
2. `PORT=8080 gradle run`
3. `grpcurl -proto app/src/main/proto/example.proto -plaintext localhost:8080 example.Example.Run`
