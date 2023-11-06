from concurrent import futures
import logging
import math
import time

import grpc
import advent_pb2_grpc

class SolverService(advent_pb2_grpc.SolverServiceServicer):
    def Solve(self, request, context):
        return None
    

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    advent_pb2_grpc.add_SolverServiceServicer_to_server(SolverService(), server)
    server.add_insecure_port("[::]:8080")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    serve()