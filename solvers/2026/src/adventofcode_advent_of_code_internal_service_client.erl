%%%-------------------------------------------------------------------
%% @doc Client module for grpc service adventofcode.AdventOfCodeInternalService.
%% @end
%%%-------------------------------------------------------------------

%% this module was generated and should not be modified manually

-module(adventofcode_advent_of_code_internal_service_client).

-compile(export_all).
-compile(nowarn_export_all).

-include_lib("grpcbox/include/grpcbox.hrl").

-define(is_ctx(Ctx), is_tuple(Ctx) andalso element(1, Ctx) =:= ctx).

-define(SERVICE, 'adventofcode.AdventOfCodeInternalService').
-define(PROTO_MODULE, 'advent_pb').
-define(MARSHAL_FUN(T), fun(I) -> ?PROTO_MODULE:encode_msg(I, T) end).
-define(UNMARSHAL_FUN(T), fun(I) -> ?PROTO_MODULE:decode_msg(I, T) end).
-define(DEF(Input, Output, MessageType), #grpcbox_def{service=?SERVICE,
                                                      message_type=MessageType,
                                                      marshal_fun=?MARSHAL_FUN(Input),
                                                      unmarshal_fun=?UNMARSHAL_FUN(Output)}).

%% @doc Unary RPC
-spec upload(advent_pb:upload_request()) ->
    {ok, advent_pb:upload_response(), grpcbox:metadata()} | grpcbox_stream:grpc_error_response() | {error, any()}.
upload(Input) ->
    upload(ctx:new(), Input, #{}).

-spec upload(ctx:t() | advent_pb:upload_request(), advent_pb:upload_request() | grpcbox_client:options()) ->
    {ok, advent_pb:upload_response(), grpcbox:metadata()} | grpcbox_stream:grpc_error_response() | {error, any()}.
upload(Ctx, Input) when ?is_ctx(Ctx) ->
    upload(Ctx, Input, #{});
upload(Input, Options) ->
    upload(ctx:new(), Input, Options).

-spec upload(ctx:t(), advent_pb:upload_request(), grpcbox_client:options()) ->
    {ok, advent_pb:upload_response(), grpcbox:metadata()} | grpcbox_stream:grpc_error_response() | {error, any()}.
upload(Ctx, Input, Options) ->
    grpcbox_client:unary(Ctx, <<"/adventofcode.AdventOfCodeInternalService/Upload">>, Input, ?DEF(upload_request, upload_response, <<"adventofcode.UploadRequest">>), Options).

%% @doc Unary RPC
-spec register(advent_pb:register_request()) ->
    {ok, advent_pb:register_response(), grpcbox:metadata()} | grpcbox_stream:grpc_error_response() | {error, any()}.
register(Input) ->
    register(ctx:new(), Input, #{}).

-spec register(ctx:t() | advent_pb:register_request(), advent_pb:register_request() | grpcbox_client:options()) ->
    {ok, advent_pb:register_response(), grpcbox:metadata()} | grpcbox_stream:grpc_error_response() | {error, any()}.
register(Ctx, Input) when ?is_ctx(Ctx) ->
    register(Ctx, Input, #{});
register(Input, Options) ->
    register(ctx:new(), Input, Options).

-spec register(ctx:t(), advent_pb:register_request(), grpcbox_client:options()) ->
    {ok, advent_pb:register_response(), grpcbox:metadata()} | grpcbox_stream:grpc_error_response() | {error, any()}.
register(Ctx, Input, Options) ->
    grpcbox_client:unary(Ctx, <<"/adventofcode.AdventOfCodeInternalService/Register">>, Input, ?DEF(register_request, register_response, <<"adventofcode.RegisterRequest">>), Options).

%% @doc Unary RPC
-spec add_solution(advent_pb:add_solution_request()) ->
    {ok, advent_pb:add_solution_response(), grpcbox:metadata()} | grpcbox_stream:grpc_error_response() | {error, any()}.
add_solution(Input) ->
    add_solution(ctx:new(), Input, #{}).

-spec add_solution(ctx:t() | advent_pb:add_solution_request(), advent_pb:add_solution_request() | grpcbox_client:options()) ->
    {ok, advent_pb:add_solution_response(), grpcbox:metadata()} | grpcbox_stream:grpc_error_response() | {error, any()}.
add_solution(Ctx, Input) when ?is_ctx(Ctx) ->
    add_solution(Ctx, Input, #{});
add_solution(Input, Options) ->
    add_solution(ctx:new(), Input, Options).

-spec add_solution(ctx:t(), advent_pb:add_solution_request(), grpcbox_client:options()) ->
    {ok, advent_pb:add_solution_response(), grpcbox:metadata()} | grpcbox_stream:grpc_error_response() | {error, any()}.
add_solution(Ctx, Input, Options) ->
    grpcbox_client:unary(Ctx, <<"/adventofcode.AdventOfCodeInternalService/AddSolution">>, Input, ?DEF(add_solution_request, add_solution_response, <<"adventofcode.AddSolutionRequest">>), Options).

%% @doc Unary RPC
-spec get_solution(advent_pb:get_solution_request()) ->
    {ok, advent_pb:get_solution_response(), grpcbox:metadata()} | grpcbox_stream:grpc_error_response() | {error, any()}.
get_solution(Input) ->
    get_solution(ctx:new(), Input, #{}).

-spec get_solution(ctx:t() | advent_pb:get_solution_request(), advent_pb:get_solution_request() | grpcbox_client:options()) ->
    {ok, advent_pb:get_solution_response(), grpcbox:metadata()} | grpcbox_stream:grpc_error_response() | {error, any()}.
get_solution(Ctx, Input) when ?is_ctx(Ctx) ->
    get_solution(Ctx, Input, #{});
get_solution(Input, Options) ->
    get_solution(ctx:new(), Input, Options).

-spec get_solution(ctx:t(), advent_pb:get_solution_request(), grpcbox_client:options()) ->
    {ok, advent_pb:get_solution_response(), grpcbox:metadata()} | grpcbox_stream:grpc_error_response() | {error, any()}.
get_solution(Ctx, Input, Options) ->
    grpcbox_client:unary(Ctx, <<"/adventofcode.AdventOfCodeInternalService/GetSolution">>, Input, ?DEF(get_solution_request, get_solution_response, <<"adventofcode.GetSolutionRequest">>), Options).

%% @doc Unary RPC
-spec set_cookie(advent_pb:set_cookie_request()) ->
    {ok, advent_pb:set_cookie_response(), grpcbox:metadata()} | grpcbox_stream:grpc_error_response() | {error, any()}.
set_cookie(Input) ->
    set_cookie(ctx:new(), Input, #{}).

-spec set_cookie(ctx:t() | advent_pb:set_cookie_request(), advent_pb:set_cookie_request() | grpcbox_client:options()) ->
    {ok, advent_pb:set_cookie_response(), grpcbox:metadata()} | grpcbox_stream:grpc_error_response() | {error, any()}.
set_cookie(Ctx, Input) when ?is_ctx(Ctx) ->
    set_cookie(Ctx, Input, #{});
set_cookie(Input, Options) ->
    set_cookie(ctx:new(), Input, Options).

-spec set_cookie(ctx:t(), advent_pb:set_cookie_request(), grpcbox_client:options()) ->
    {ok, advent_pb:set_cookie_response(), grpcbox:metadata()} | grpcbox_stream:grpc_error_response() | {error, any()}.
set_cookie(Ctx, Input, Options) ->
    grpcbox_client:unary(Ctx, <<"/adventofcode.AdventOfCodeInternalService/SetCookie">>, Input, ?DEF(set_cookie_request, set_cookie_response, <<"adventofcode.SetCookieRequest">>), Options).

%% @doc Unary RPC
-spec get_solvers(advent_pb:get_solvers_request()) ->
    {ok, advent_pb:get_solvers_response(), grpcbox:metadata()} | grpcbox_stream:grpc_error_response() | {error, any()}.
get_solvers(Input) ->
    get_solvers(ctx:new(), Input, #{}).

-spec get_solvers(ctx:t() | advent_pb:get_solvers_request(), advent_pb:get_solvers_request() | grpcbox_client:options()) ->
    {ok, advent_pb:get_solvers_response(), grpcbox:metadata()} | grpcbox_stream:grpc_error_response() | {error, any()}.
get_solvers(Ctx, Input) when ?is_ctx(Ctx) ->
    get_solvers(Ctx, Input, #{});
get_solvers(Input, Options) ->
    get_solvers(ctx:new(), Input, Options).

-spec get_solvers(ctx:t(), advent_pb:get_solvers_request(), grpcbox_client:options()) ->
    {ok, advent_pb:get_solvers_response(), grpcbox:metadata()} | grpcbox_stream:grpc_error_response() | {error, any()}.
get_solvers(Ctx, Input, Options) ->
    grpcbox_client:unary(Ctx, <<"/adventofcode.AdventOfCodeInternalService/GetSolvers">>, Input, ?DEF(get_solvers_request, get_solvers_response, <<"adventofcode.GetSolversRequest">>), Options).

