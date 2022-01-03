# H2C & GRPC Sample Apps

Cloud Foundry sample apps using plain text HTTP2 (H2C) for Go, Node.JS, Ruby, and Python.

Note: due to the lack of H2C support in popular frameworks, the Ruby and Python examples are very low level.

### How to deploy HTTP2 / H2C apps

Make sure you have [CF CLI](https://docs.cloudfoundry.org/cf-cli/install-go-cli.html) version 8 or higher installed

```shell
> cf version
cf version 8.0.0+e8d41cf8e.2021-09-16
```

Check out this repo and go to the app you want to test

```shell
> git clone https://github.com/SAP-samples/cf-http2
> cd go-http2
```

choose a name and push the application without starting

```shell
> cf push --no-route http2-example-app-go
```

Map an HTTP2 route with the application

```shell
> export CF_APPS_DOMAIN=my-apps.cf.example.com
> cf map-route http2-example-app-go $CF_APPS_DOMAIN --hostname http2-example-app-go --destination-protocol http2
```

Check the app is working

```shell
> curl https://http2-example-app-go.$CF_APPS_DOMAIN /
Hello! This Go application is speaking plain text HTTP2 (H2C) with the CF routing layer
```

### How to deploy gRPC apps

Make sure you have [CF CLI](https://docs.cloudfoundry.org/cf-cli/install-go-cli.html) version 8 or higher installed

```shell
> cf version
cf version 8.0.0+e8d41cf8e.2021-09-16
```

Check out this repo and go to the app you want to test

```shell
> git clone https://github.com/SAP-samples/cf-gRPC
> cd go-gRPC
```

Choose a name and push the application without starting

```shell
> cf push --no-route gRPC-example-app-go
```

Map an HTTP2 route with the application. It must be a route on an apps domain supporting mutual TLS (mTLS) and you must have the client certificate and private key available on your machine

```shell
> export CF_MTLS_APPS_DOMAIN=my-mtls-apps.cf.example.com
> cf map-route grpc-example-app-go $CF_MTLS_APPS_DOMAIN --hostname grpc-example-app-go --destination-protocol http2
```

Download grpcurl from https://github.com/fullstorydev/grpcurl/releases and put it in your path

FIXME:
```
> export MTLS_CERT_PATH=path/to/mtls_client_certificate.pem
> export MTLS_KEY_PATH=path/to/mtls_client_private_key.pem
> grpcurl -insecure -cert $MTLS_CERT_PATH" -key $MTLS_KEY_PATH" grpc-example-app-go.$CF_MTLS_APPS_DOMAIN:443 example.Example.Run
```
