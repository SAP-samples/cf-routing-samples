// based on https://github.com/grpc/grpc/blob/master/examples/node/dynamic_codegen/greeter_server.js

const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const packageDefinition = protoLoader.loadSync(__dirname + '/example.proto');
const exampleProto = grpc.loadPackageDefinition(packageDefinition);

function run(_, callback) {
  callback(null, {message: 'Hello! This Node.JS application is speaking gRPC'});
}

const port = process.env.PORT;
const server = new grpc.Server();

server.addService(exampleProto.Example.service, { run: run });
server.bindAsync(`0.0.0.0:${port}`, grpc.ServerCredentials.createInsecure(), () => {
  server.start();
});
