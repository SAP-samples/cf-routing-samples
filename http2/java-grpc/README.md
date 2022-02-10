# Java gRPC Sample

Example gRPC server implemented in Java
## Running Locally

1. `./gradlew build`
2. `PORT=8080 ./gradlew run`
3. `grpcurl -proto app/src/main/proto/example.proto -plaintext localhost:8080 example.Example.Run`

This example uses Gradle Wrapper, which will download and use the appropriate Gradle version to build the project.

A Java installation, e.g. openjdk, is required for this sample.
