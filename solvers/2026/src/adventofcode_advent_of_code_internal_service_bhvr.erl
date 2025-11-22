%%%-------------------------------------------------------------------
%% @doc Behaviour to implement for grpc service adventofcode.AdventOfCodeInternalService.
%% @end
%%%-------------------------------------------------------------------

%% this module was generated and should not be modified manually

-module(adventofcode_advent_of_code_internal_service_bhvr).

%% Unary RPC
-callback upload(ctx:t(), advent_pb:upload_request()) ->
    {ok, advent_pb:upload_response(), ctx:t()} | grpcbox_stream:grpc_error_response().

%% Unary RPC
-callback register(ctx:t(), advent_pb:register_request()) ->
    {ok, advent_pb:register_response(), ctx:t()} | grpcbox_stream:grpc_error_response().

%% Unary RPC
-callback add_solution(ctx:t(), advent_pb:add_solution_request()) ->
    {ok, advent_pb:add_solution_response(), ctx:t()} | grpcbox_stream:grpc_error_response().

%% Unary RPC
-callback get_solution(ctx:t(), advent_pb:get_solution_request()) ->
    {ok, advent_pb:get_solution_response(), ctx:t()} | grpcbox_stream:grpc_error_response().

%% Unary RPC
-callback set_cookie(ctx:t(), advent_pb:set_cookie_request()) ->
    {ok, advent_pb:set_cookie_response(), ctx:t()} | grpcbox_stream:grpc_error_response().

%% Unary RPC
-callback get_solvers(ctx:t(), advent_pb:get_solvers_request()) ->
    {ok, advent_pb:get_solvers_response(), ctx:t()} | grpcbox_stream:grpc_error_response().

