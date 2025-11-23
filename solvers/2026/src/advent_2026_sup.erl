-module(advent_2026_sup).

-behaviour(supervisor).

-export([start_link/0]).

-export([init/1]).

-define(SERVER, ?MODULE).

start_link() ->
    supervisor:start_link({local, ?SERVER}, ?MODULE, []).

init([]) ->
    SupFlags = #{strategy => one_for_all,
                 intensity => 0,
                 period => 1},
    ChildSpecs = [
        #{id => advent_2026_register,
          start => {advent_2026_register, start_link, []},
          restart => permanent,
          shutdown => 5000,
          type => worker,
          modules => [advent_2026_register]}
    ],
    {ok, {SupFlags, ChildSpecs}}.
