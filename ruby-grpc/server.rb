require_relative './example_services_pb.rb'
require_relative './example_pb.rb'

class ExampleServer < Example::Service
  def run(_, _)
    Response::new(message: "Hello! This Ruby application is speaking gRPC")
  end
end

port = ENV.fetch("PORT")

server = GRPC::RpcServer.new
server.add_http2_port("0.0.0.0:#{port}", :this_port_is_insecure)
server.handle(ExampleServer)
server.run_till_terminated_or_interrupted([1, 'int', 'SIGQUIT'])
