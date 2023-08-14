# Java gRPC Sample

Example gRPC server implemented in Java
## Deploying the Java app
### Pre-requisite for Java app
Please follow [the install setup for Java](../README.md#for-java-app-do-the-following) first before deploying this java app.

### Deploy the app using the manifest file

Deploy the app
```shell
export DOMAIN=my.cf.app.domain
cf push -f app-manifest.yml --var domain=$DOMAIN
```
### Deploy the app without the manifest file
```shell
cf push --no-route java-grpc-test --buildpack java_buildpack --path ./app/build/distributions/app.zip
# my.cf.app.domain is used as an example for demonstration purpose
cf map-route java-grpc-test my.cf.app.domain --hostname java-grpc-test --app-protocol http2
```

## Testing the Java app
`grpcurl` needs to be [installed separately](https://github.com/fullstorydev/grpcurl).
```shell
grpcurl -proto app/src/main/proto/example.proto java-grpc-test.my.cf.app.domain:443 example.Example.Run 
```
## Building and running on the local machine

1. `./gradlew build`
2. `PORT=8080 ./gradlew run`
3. `grpcurl -proto app/src/main/proto/example.proto -plaintext localhost:8080 example.Example.Run`

This example uses Gradle Wrapper, which will download and use the appropriate Gradle version to build the project.

A Java installation, e.g. openjdk, is required for this sample.
