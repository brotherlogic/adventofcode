-module(aoc_server_sup).
-behaviour(supervisor).

-export([start_link/0, init/1]).

start_link() ->
    supervisor:start_link({local, ?MODULE}, ?MODULE, []).

init([]) ->
    Port = 8080,
    Services = #{adventofcode_solver_service_pb => aoc_server_grpc},
    Children = [
        #{
            id => grpc_server,
            start => {grpc, start_server, [Port, Services, []]},
            restart => permanent,
            shutdown => 5000,
            type => worker,
            modules => [aoc_server_grpc]
        }
    ],
    SupFlags = #{strategy => one_for_one, intensity => 1, period => 5},
    {ok, {SupFlags, Children}}.
