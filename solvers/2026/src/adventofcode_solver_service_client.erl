%%%-------------------------------------------------------------------
%% @doc Client module for grpc service adventofcode.SolverService.
%% @end
%%%-------------------------------------------------------------------

%% this module was generated and should not be modified manually

-module(adventofcode_solver_service_client).

-compile(export_all).
-compile(nowarn_export_all).

-include_lib("grpcbox/include/grpcbox.hrl").

-define(is_ctx(Ctx), is_tuple(Ctx) andalso element(1, Ctx) =:= ctx).

-define(SERVICE, 'adventofcode.SolverService').
-define(PROTO_MODULE, 'advent_pb').
-define(MARSHAL_FUN(T), fun(I) -> ?PROTO_MODULE:encode_msg(I, T) end).
-define(UNMARSHAL_FUN(T), fun(I) -> ?PROTO_MODULE:decode_msg(I, T) end).
-define(DEF(Input, Output, MessageType), #grpcbox_def{service=?SERVICE,
                                                      message_type=MessageType,
                                                      marshal_fun=?MARSHAL_FUN(Input),
                                                      unmarshal_fun=?UNMARSHAL_FUN(Output)}).

%% @doc Unary RPC
-spec solve(advent_pb:solve_request()) ->
    {ok, advent_pb:solve_response(), grpcbox:metadata()} | grpcbox_stream:grpc_error_response() | {error, any()}.
solve(Input) ->
    solve(ctx:new(), Input, #{}).

-spec solve(ctx:t() | advent_pb:solve_request(), advent_pb:solve_request() | grpcbox_client:options()) ->
    {ok, advent_pb:solve_response(), grpcbox:metadata()} | grpcbox_stream:grpc_error_response() | {error, any()}.
solve(Ctx, Input) when ?is_ctx(Ctx) ->
    solve(Ctx, Input, #{});
solve(Input, Options) ->
    solve(ctx:new(), Input, Options).

-spec solve(ctx:t(), advent_pb:solve_request(), grpcbox_client:options()) ->
    {ok, advent_pb:solve_response(), grpcbox:metadata()} | grpcbox_stream:grpc_error_response() | {error, any()}.
solve(Ctx, Input, Options) ->
    grpcbox_client:unary(Ctx, <<"/adventofcode.SolverService/Solve">>, Input, ?DEF(solve_request, solve_response, <<"adventofcode.SolveRequest">>), Options).

