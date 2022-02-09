from concurrent import futures
import logging
import os

import grpc
from grpc_reflection.v1alpha import reflection
import example_pb2
import example_pb2_grpc


class Example(example_pb2_grpc.ExampleServicer):
    def Run(self, _request, _context):
        return example_pb2.Response(message='Hello! This Python application is speaking gRPC')


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    example_pb2_grpc.add_ExampleServicer_to_server(Example(), server)
    SERVICE_NAMES = (
        example_pb2.DESCRIPTOR.services_by_name['Example'].full_name,
        reflection.SERVICE_NAME,
    )
    reflection.enable_server_reflection(SERVICE_NAMES, server)

    port = os.getenv('PORT')
    server.add_insecure_port("[::]:%s" %(port))
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
