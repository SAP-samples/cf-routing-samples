

const http2 = require('http2');
const port = process.env.PORT;

http2.createServer( (_, res) => {
  res.writeHead(200, {'Content-Type': 'text/plain'});
  res.end("Hello! This Node.js application is speaking HTTP2 with the CF routing layer");
}).listen(port, () => {
  console.log("Listening on " + port);
});
