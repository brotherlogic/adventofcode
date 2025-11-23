-module(advent_2026_register).
-behaviour(gen_server).

%% API
-export([start_link/0]).

%% gen_server callbacks
-export([init/1, handle_call/3, handle_cast/2, handle_info/2, terminate/2,
         code_change/3]).

-define(SERVER, ?MODULE).
-define(INTERVAL, 60000). % 1 minute

%%%===================================================================
%%% API
%%%===================================================================

start_link() ->
    gen_server:start_link({local, ?SERVER}, ?MODULE, [], []).

%%%===================================================================
%%% gen_server callbacks
%%%===================================================================

init([]) ->
    self() ! register,
    {ok, []}.

handle_call(_Request, _From, State) ->
    {reply, ok, State}.

handle_cast(_Msg, State) ->
    {noreply, State}.

handle_info(register, State) ->
    register(),
    erlang:send_after(?INTERVAL, self(), register),
    {noreply, State};
handle_info(_Info, State) ->
    {noreply, State}.

terminate(_Reason, _State) ->
    ok.

code_change(_OldVsn, State, _Extra) ->
    {ok, State}.

%%%===================================================================
%%% Internal functions
%%%===================================================================

register() ->
    Request = #{
        callback => <<"adventofcode-solver-2026.adventofcode:8080">>,
        year => 2026
    },
    Ctx = ctx:new(),
    % Note: We assume the channel is configured or we pass it directly.
    % Since we don't have sys.config setup for this specific peer, we might need to add it there
    % or try to pass it as an option if grpcbox supports it.
    % For now, I will try to assume the environment might have it or I'll add it to sys.config.
    % We use the default channel configured in sys.config
    Options = #{},
    case adventofcode_advent_of_code_internal_service_client:register(Ctx, Request, Options) of
        {ok, _Response, _Metadata} ->
            io:format("Registered successfully~n");
        Error ->
            io:format("Failed to register: ~p~n", [Error])
    end.
