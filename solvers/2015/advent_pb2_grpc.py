# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import advent_pb2 as advent__pb2


class AdventOfCodeServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Solve = channel.unary_unary(
                '/adventofcode.AdventOfCodeService/Solve',
                request_serializer=advent__pb2.SolveRequest.SerializeToString,
                response_deserializer=advent__pb2.SolveResponse.FromString,
                )
        self.Upload = channel.unary_unary(
                '/adventofcode.AdventOfCodeService/Upload',
                request_serializer=advent__pb2.UploadRequest.SerializeToString,
                response_deserializer=advent__pb2.UploadResponse.FromString,
                )
        self.Register = channel.unary_unary(
                '/adventofcode.AdventOfCodeService/Register',
                request_serializer=advent__pb2.RegisterRequest.SerializeToString,
                response_deserializer=advent__pb2.RegisterResponse.FromString,
                )


class AdventOfCodeServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Solve(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Upload(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Register(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AdventOfCodeServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Solve': grpc.unary_unary_rpc_method_handler(
                    servicer.Solve,
                    request_deserializer=advent__pb2.SolveRequest.FromString,
                    response_serializer=advent__pb2.SolveResponse.SerializeToString,
            ),
            'Upload': grpc.unary_unary_rpc_method_handler(
                    servicer.Upload,
                    request_deserializer=advent__pb2.UploadRequest.FromString,
                    response_serializer=advent__pb2.UploadResponse.SerializeToString,
            ),
            'Register': grpc.unary_unary_rpc_method_handler(
                    servicer.Register,
                    request_deserializer=advent__pb2.RegisterRequest.FromString,
                    response_serializer=advent__pb2.RegisterResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'adventofcode.AdventOfCodeService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class AdventOfCodeService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Solve(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/adventofcode.AdventOfCodeService/Solve',
            advent__pb2.SolveRequest.SerializeToString,
            advent__pb2.SolveResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Upload(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/adventofcode.AdventOfCodeService/Upload',
            advent__pb2.UploadRequest.SerializeToString,
            advent__pb2.UploadResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Register(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/adventofcode.AdventOfCodeService/Register',
            advent__pb2.RegisterRequest.SerializeToString,
            advent__pb2.RegisterResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)


class SolverServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Solve = channel.unary_unary(
                '/adventofcode.SolverService/Solve',
                request_serializer=advent__pb2.SolveRequest.SerializeToString,
                response_deserializer=advent__pb2.SolveResponse.FromString,
                )


class SolverServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Solve(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SolverServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Solve': grpc.unary_unary_rpc_method_handler(
                    servicer.Solve,
                    request_deserializer=advent__pb2.SolveRequest.FromString,
                    response_serializer=advent__pb2.SolveResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'adventofcode.SolverService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class SolverService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Solve(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/adventofcode.SolverService/Solve',
            advent__pb2.SolveRequest.SerializeToString,
            advent__pb2.SolveResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
