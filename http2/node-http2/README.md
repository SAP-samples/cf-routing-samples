# Node.JS HTTP/2 H2C Sample

Example HTTP/2 server implemented in JavaScript / Node.js

The app supports HTTP/1.1 and HTTP/2 H2C if requested explicitly, i.e. no connection upgrade.

## Running Locally

1. `npm install`
2. `PORT=8080 npm start`
3. `curl -v --http2-prior-knowledge http://localhost:8080`
