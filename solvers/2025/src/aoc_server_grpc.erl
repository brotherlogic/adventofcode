-module(aoc_server_grpc).
-behaviour(adventofcode_solver_service_pb).

-include("adventofcode_pb.hrl").

-export([solve/2]).

-record(state, {}).

solve(Req, _Stream) ->
    io:format("Received solve request: ~p~n", [Req]),
    Response = #'SolveResponse'{
        answer = 2025,
        string_answer = "hello world",
        big_answer = 12
    },
    {Response, _Stream}.