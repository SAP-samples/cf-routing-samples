# gRPC Example App

Example gRPC server implemented in Python

## Running Locally FIXME
1. `pipenv install -r requirements.txt`
2. `PORT=8080 python3 main.py`

### Renerating code

```shell
python -m pip install grpcio-tools
python -m grpc_tools.protoc -I . --python_out . --grpc_python_out . example.proto
```
