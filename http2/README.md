# H2C & gRPC Sample Apps

This folder contains Cloud Foundry sample apps showcasing plain text HTTP2 (H2C) and gRPC via secure HTTP2.

A more detailed look into HTTP/2 on the SAP BTP, Cloud Foundry environment can be found in the following blog post:
* [HTTP/2 on the SAP BTP, Cloud Foundry runtime](https://blogs.sap.com/2022/02/16/http-2-on-sap-btp-cloud-foundry-runtime/)

Each sample has a README file that explains how to run it locally. Some of these examples require additional applications or commands that can be obtained from their respective home pages or operating system distribution.

Note: due to the lack of H2C support in popular frameworks, the Ruby and Python examples are very low level.

## Downloading the Samples

Clone this repo 

```shell
# Skip this if you cloned the repository already
git clone https://github.com/SAP-samples/cf-routing-samples
cd cf-routing-samples/http2
```

## Deploying all Apps

Each app can be deployed individually, or all apps can be deployed at the same time using a common manifest provided in [apps-manifest.yml](apps-manifest.yml).

The following sections explain pre-requisites and instructions for the different sample apps.

### Pre-requisites
Make sure you have [CF CLI](https://docs.cloudfoundry.org/cf-cli/install-go-cli.html) version 8 or higher installed

```shell
cf version
```
example output:
```plain
cf version 8.0.0+e8d41cf8e.2021-09-16
```

#### Pre-requisites for Java apps

```shell
# please make sure that the openjdk@11 is installed and gradle wrapper is pointing to it
./gradlew clean build
```

### Deploying Apps Individually

#### HTTP/2 Plain Text (H2C)
Please check the README file for each of the HTTP/2 (H2C) example apps:
* Go: [go-http2](go-http2)
* Java: [java-http2](java-http2) (Note: refer [pre-requisites](#pre-requisites-for-java-apps) for Java)
* Node: [node-http2](node-http2)
* Python: [python-http2](python-http2)
* Ruby: [ruby-http2](ruby-http2)

#### gRPC
Please check the README file for each of the gRPC example apps:
* Go: [go-grpc](go-grpc)
* Java: [java-grpc](java-grpc) (Note: refer [pre-req](#For java app do the following) for Java)
* Node: [node-grpc](node-grpc)
* Python: [python-grpc](python-grpc)
* Ruby: [ruby-grpc](ruby-grpc)

### Deploying all Apps with one Manifest
```shell
# run the following commands in cf environment
cd http2
export DOMAIN=my.cf.app.domain
cf push --manifest apps-manifest.yml --var domain=$DOMAIN --vars-file gradle.properties
```

### Testing Sample Apps

#### Testing H2C Apps
```shell
# format: https://<language>-http2-test.my.cf.app.domain
export DOMAIN=my.cf.app.domain
curl -v --http2-prior-knowledge "https://go-http2-test.$DOMAIN"
curl -v --http2-prior-knowledge "https://java-http2-test.$DOMAIN"
curl -v --http2-prior-knowledge "https://node-http2-test.$DOMAIN"
curl -v --http2-prior-knowledge "https://python-http2-test.$DOMAIN"
curl -v --http2-prior-knowledge "https://ruby-http2-test.$DOMAIN"
```

The output of each of those commands should be HTTP 200 with the response body that reflects the app's implementation language.

Please note that for H2C, `--http2-prior-knowledge` is mandatory, as without [TLS](https://en.wikipedia.org/wiki/Transport_Layer_Security) there is no [ALPN](https://en.wikipedia.org/wiki/Application-Layer_Protocol_Negotiation).

#### Testing gRPC Apps

1. Download grpcurl from https://github.com/fullstorydev/grpcurl/releases and put it in your path
2. Make a gRPC request using grpcurl
   ```shell
   export DOMAIN=my.cf.app.domain
   # following example with mtls
   export MTLS_CERT_PATH=path/to/mtls_client_certificate.pem
   export MTLS_KEY_PATH=path/to/mtls_client_private_key.pem
   
   grpcurl -insecure -cert $MTLS_CERT_PATH -key $MTLS_KEY_PATH "go-grpc-test.$DOMAIN:443" example.Example.Run
   grpcurl -insecure -cert $MTLS_CERT_PATH -key $MTLS_KEY_PATH "java-grpc-test.$DOMAIN:443" example.Example.Run
   grpcurl -insecure -cert $MTLS_CERT_PATH -key $MTLS_KEY_PATH "python-grpc-test.$DOMAIN:443" example.Example.Run
   grpcurl -insecure -cert $MTLS_CERT_PATH -key $MTLS_KEY_PATH "ruby-grpc-test.$DOMAIN:443" example.Example.Run
   
   # Alternative grpcurl for apps which don't support reflection (such as Node.js)
   grpcurl -proto path/to/example.proto -insecure -cert "$MTLS_CERT_PATH" -key "$MTLS_KEY_PATH" "node-grpc-test.$DOMAIN:443" Example.Run
   ```

The output of each of those commands should be in the following format, with a message that reflects the app's implementation language:
```json
{
  "message": "Hello! This Go application is speaking gRPC"
}
```