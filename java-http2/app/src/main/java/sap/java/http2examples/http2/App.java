package sap.java.http2examples.http2;

import reactor.core.publisher.Mono;
import reactor.netty.DisposableServer;
import reactor.netty.http.HttpProtocol;
import reactor.netty.http.server.HttpServer;

public class App {

	public static void main(String[] args) {
        int port = Integer.parseInt(System.getenv("PORT"));

		DisposableServer server =
				HttpServer.create()
				          .port(port)
				          .protocol(HttpProtocol.H2C)
				          .handle((request, response) -> response.sendString(Mono.just("Hello! This Java application is speaking plain text HTTP2 (H2C) with the CF routing layer\n")))
				          .bindNow();

		server.onDispose()
		      .block();
	}
}
