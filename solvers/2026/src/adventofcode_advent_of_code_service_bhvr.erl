%%%-------------------------------------------------------------------
%% @doc Behaviour to implement for grpc service adventofcode.AdventOfCodeService.
%% @end
%%%-------------------------------------------------------------------

%% this module was generated and should not be modified manually

-module(adventofcode_advent_of_code_service_bhvr).

%% Unary RPC
-callback solve(ctx:t(), advent_pb:solve_request()) ->
    {ok, advent_pb:solve_response(), ctx:t()} | grpcbox_stream:grpc_error_response().

