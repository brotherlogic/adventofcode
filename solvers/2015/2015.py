from concurrent import futures
import logging
import math
import time

import grpc
import advent_pb2
import advent_pb2_grpc

import day1

class SolverService(advent_pb2_grpc.SolverServiceServicer):
    def Solve(self, request, context):
        if request.year == 2015 and request.day == 1 and request.part == 1:
            return day1.SolveDay1Part1(request.data)

        return None

def register():
    channel = grpc.insecure_channel('adventofcode.adventofcode:8082')
    stub = advent_pb2_grpc.AdventOfCodeServiceStub(channel)
    response = stub.Register(advent_pb2.RegisterRequest(year=2015, callback="adventofcode.adventofcode-solver-2015:8080"))

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    advent_pb2_grpc.add_SolverServiceServicer_to_server(SolverService(), server)
    server.add_insecure_port("[::]:8080")
    server.start()
    register()
    server.wait_for_termination()

if __name__ == "__main__":
    serve()