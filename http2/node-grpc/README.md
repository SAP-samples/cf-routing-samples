# Node.js gRPC Sample

Example gRPC server implemented in Node.JS
## Deploying the Node app
### Deploy the app using the manifest file
Deploy the app
```shell
export DOMAIN=my.cf.app.domain
cf push -f app-manifest.yml --var domain=$DOMAIN
```
### Deploy the app without the manifest file
```shell
cf push --no-route node-grpc-test --buildpack https://github.com/cloudfoundry/nodejs-buildpack 
# my.cf.app.domain is used as an example for demonstration purpose
cf map-route node-grpc-test my.cf.app.domain --hostname node-grpc-test --app-protocol http2
```

## Testing the Node app
`grpcurl` needs to be [installed separately](https://github.com/fullstorydev/grpcurl).
```shell
grpcurl -proto example.proto node-grpc-test.my.cf.app.domain:443 Example.Run 
```

## Building and running on the local machine

1. `npm install`
2. `PORT=8080 npm start`
3. `grpcurl -proto example.proto -plaintext localhost:8080 Example.Run`
