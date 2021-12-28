# H2C Sample Apps

Cloud Foundry sample apps using plain-text HTTP2 (H2C)

### How to deploy

Make sure you have [CF CLI](https://docs.cloudfoundry.org/cf-cli/install-go-cli.html) version 8 or higher installed

```
cf version
```

Check out this repo and go to the app you want to test

```
git clone https://github.com/SAP-samples/cf-http2
cd go
```

choose a name and push the application without starting

```
cf push --no-route http2-example-app-go
```

Map an HTTP2 route with the application

```
cf map-route http2-example-app-go cfapps.cfi01.ali.cfi.sapcloud.io --hostname http2-example-app-go --destination-protocol http2
```

Check the app is working

```
curl https://http2-example-app-go.cfapps.cfi01.ali.cfi.sapcloud.io/
```
