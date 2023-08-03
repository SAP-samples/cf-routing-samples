// Based on https://github.com/grpc/grpc-java/tree/master/examples/src/test/java/io/grpc/examples/helloworld
package sap.java.http2examples.grpc;

import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.stub.StreamObserver;
import java.io.IOException;
import java.util.concurrent.TimeUnit;
import java.util.logging.Logger;


public class App {
  private static final Logger logger = Logger.getLogger(App.class.getName());

  private Server server;

  private void start() throws IOException {
    int port = Integer.parseInt(System.getenv("PORT"));
    server = ServerBuilder.forPort(port)
        .addService(new ExampleImpl())
        .build()
        .start();
    logger.info("Server started, listening on " + port);
    Runtime.getRuntime().addShutdownHook(new Thread() {
      @Override
      public void run() {
        // Use stderr here since the logger may have been reset by its JVM shutdown hook.
        System.err.println("*** shutting down gRPC server since JVM is shutting down");
        try {
          App.this.stop();
        } catch (InterruptedException e) {
          e.printStackTrace(System.err);
        }
        System.err.println("*** server shut down");
      }
    });
  }

  private void stop() throws InterruptedException {
    if (server != null) {
      server.shutdown().awaitTermination(30, TimeUnit.SECONDS);
    }
  }

  /**
   * Await termination on the main thread since the grpc library uses daemon threads.
   */
  private void blockUntilShutdown() throws InterruptedException {
    if (server != null) {
      server.awaitTermination();
    }
  }

  /**
   * Main launches the server from the command line.
   */
  public static void main(String[] args) throws IOException, InterruptedException {
    final App server = new App();
    server.start();
    server.blockUntilShutdown();
  }

  static class ExampleImpl extends ExampleGrpc.ExampleImplBase {
    @Override
    public void run(Request req, StreamObserver<Response> responseObserver) {
      Response response = Response.newBuilder().setMessage("Hello! This Java application is speaking gRPC").build();
      responseObserver.onNext(response);
      responseObserver.onCompleted();
    }
  }
}
