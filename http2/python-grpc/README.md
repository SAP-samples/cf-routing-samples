# Python gRPC Sample

Example gRPC server implemented in Python
## Deploying the Python app
### Deploy the app using the manifest file
Deploy the app
```shell
> export DOMAIN=my.cf.app.domain
> cf push -f app-manifest.yml --var domain=$DOMAIN
```

### Deploy the app without the manifest file
```shell
> cf push --no-route python-grpc-test --buildpack https://github.com/cloudfoundry/python-buildpack 
# my.cf.app.domain is used as an example for demonstration purpose
> cf map-route python-grpc-test my.cf.app.domain --hostname python-grpc-test --app-protocol http2
```

## Testing the Python app
`grpcurl` needs to be [installed separately](https://github.com/fullstorydev/grpcurl).
```shell
> grpcurl -proto example.proto python-grpc-test.my.cf.app.domain:443 Example.Run
```

## Building and running on the local machine
1. `pipenv install -r requirements.txt`
2. `PORT=8080 pipenv run python3 main.py`
3. `grpcurl -proto example.proto -plaintext localhost:8080 Example.Run`

### Renerating code

```shell
python -m pip install grpcio-tools
python -m grpc_tools.protoc -I . --python_out . --grpc_python_out . example.proto
```
