# HTTP/2 H2C Sample

Example HTTP/2 server implemented in Go
## Deploying the go app
### Deploy the app using the manifest file

Deploy the app
```shell
export DOMAIN=my.cf.app.domain
cf push -f app-manifest.yml --var domain=$DOMAIN
```

### Deploy the app without the manifest file
```shell
cf push --no-route go-http2-test --buildpack go_buildpack
# my.cf.app.domain is used as an example for demonstration purpose
cf map-route go-http2-test my.cf.app.domain --hostname go-http2-test --app-protocol http2
```

## Testing the Go app
```shell
curl -v --http2-prior-knowledge https://go-http2-test.my.cf.app.domain 
```

The app supports HTTP/1.1 and HTTP/2 H2C if requested explicitly, i.e. no connection upgrade.

## Building and running on the local machine

1. `go build`
2. `PORT=8080 ./go-http2`
3. `curl -v --http2-prior-knowledge http://localhost:8080`
