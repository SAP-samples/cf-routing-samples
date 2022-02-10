import socket
import os

import h2.connection
import h2.config

# Example H2C web app
# Based on https://python-hyper.org/projects/h2/en/stable/basic-usage.html#writing-your-server


def handle(sock):
    config = h2.config.H2Configuration(client_side=False)
    conn = h2.connection.H2Connection(config=config)
    conn.initiate_connection()
    sock.sendall(conn.data_to_send())

    while True:
        data = sock.recv(65535)
        if not data:
            break

        events = conn.receive_data(data)
        for event in events:
            if isinstance(event, h2.events.RequestReceived):
                stream_id = event.stream_id
                conn.send_headers(
                    stream_id=stream_id,
                    headers=[
                        (':status', '200'),
                        ('content-type', 'text/plain')
                    ],
                )
                conn.send_data(
                    stream_id=stream_id,
                    data=b'Hello! This Python application is speaking plain text HTTP2 (H2C) with the CF routing layer\n',
                    end_stream=True
                )

        data_to_send = conn.data_to_send()
        if data_to_send:
            sock.sendall(data_to_send)

sock = socket.socket()
sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
sock.bind(('0.0.0.0', int(os.getenv('PORT'))))
sock.listen(5)

while True:
    handle(sock.accept()[0])
