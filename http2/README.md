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
> git clone https://github.com/SAP-samples/cf-routing-samples
> cd cf-routing-samples/http2

```
## Deploying all the apps at once
### Pre-requisites
Set the desired domain in [vars.yml](vars.yml) file:
```shell
# edit the vars.yml file and replace <domain.for.app.routes> with required domain
# my.cf.app.domain is used as an example for demonstration
> vi vars.yml
> cat vars.yml
domain: my.cf.app.domain
```

#### For java app do the following
```shell
# please make sure that the openjdk@11 is installed and gradle wrapper is pointing to it
> ./gradlew clean build
```
### Deploy all apps
```shell
# run the following commands in cf environment
> cd http2
> export DOMAIN=my.cf.app.domain
> cf push --manifest apps-manifest.yml --var domain=$DOMAIN --vars-file gradle.properties
```
## Deploying apps individually 
* Make sure you have [CF CLI](https://docs.cloudfoundry.org/cf-cli/install-go-cli.html) version 8 or higher installed

```shell
> cf version
cf version 8.0.0+e8d41cf8e.2021-09-16
```
* And set the desired domain in [vars.yml](vars.yml) file
### How to deploy H2C apps
* Go: [go-http2](go-http2)
* Java: [java-http2](java-http2) (Note: refer [pre-requisites](#For java app do the following) for Java)
* Node: [node-http2](node-http2)
* Python: [python-http2](python-http2)
* Ruby: [ruby-http2](ruby-http2)
#### Check H2C apps
```shell
# format: https://<language>-http2-test.my.cf.app.domain
> export DOMAIN=my.cf.app.domain
> curl -v --http2-prior-knowledge https://go-http2-test.my.cf.app.$DOMAIN
> curl -v --http2-prior-knowledge https://java-http2-test.my.cf.app.$DOMAIN
> curl -v --http2-prior-knowledge https://node-http2-test.my.cf.app.$DOMAIN
> curl -v --http2-prior-knowledge https://python-http2-test.my.cf.app.$DOMAIN
> curl -v --http2-prior-knowledge https://ruby-http2-test.my.cf.app.$DOMAIN
```

### How to deploy gRPC apps
* Go: [go-grpc](go-grpc)
* Java: [java-grpc](java-grpc) (Note: refer [pre-req](#For java app do the following) for Java)
* Node: [node-grpc](node-grpc)
* Python: [python-grpc](python-grpc)
* Ruby: [ruby-grpc](ruby-grpc)
#### Check gRPC apps
Download grpcurl from https://github.com/fullstorydev/grpcurl/releases and put it in your path
Make a gRPC request using grpcurl
```shell
# following example with mtls
> export MTLS_CERT_PATH=path/to/mtls_client_certificate.pem
> export MTLS_KEY_PATH=path/to/mtls_client_private_key.pem
> grpcurl -insecure -cert $MTLS_CERT_PATH -key $MTLS_KEY_PATH go-grpc-test.my.cf.app.domain:443 example.Example.Run
{
  "message": "Hello! This Go application is speaking gRPC"
}

# Alternative grpcurl for apps which don't support reflection (such as Node.js)
> grpcurl -proto path/to/example.proto -insecure -cert $MTLS_CERT_PATH -key $MTLS_KEY_PATH node-grpc-test.my.cf.app.domain:443 Example.Run
{
  "message": "Hello! This Node.JS application is speaking gRPC"
}
```
