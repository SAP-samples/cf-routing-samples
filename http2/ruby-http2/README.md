# Python HTTP/2 H2C Sample

Example HTTP/2 server implemented in Python

The app supports HTTP/1.1 and HTTP/2 H2C if requested explicitly, i.e. no connection upgrade.

## Running Locally

1. `bundle install --path vendor/bundle`
2. `PORT=8080 bundle exec ruby server.rb`
3. `curl -v --http2-prior-knowledge http://localhost:8080`
