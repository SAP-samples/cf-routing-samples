# Python gRPC Sample

Example gRPC server implemented in Python

## Running Locally
1. `pipenv install -r requirements.txt`
2. `PORT=8080 pipenv run python3 main.py`
3. `grpcurl -proto example.proto -plaintext localhost:8080 Example.Run`

### Renerating code

```shell
python -m pip install grpcio-tools
python -m grpc_tools.protoc -I . --python_out . --grpc_python_out . example.proto
```
