# Python HTTP/2 H2C Sample

Example HTTP/2 server implemented in Python

The app supports HTTP/1.1 and HTTP/2 H2C if requested explicitly, i.e. no connection upgrade.

## Running Locally

1. `pipenv install -r requirements.txt`
2. `PORT=8080 pipenv run python3 main.py`
3. `curl -v --http2-prior-knowledge http://localhost:8080`
