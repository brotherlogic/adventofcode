-module(advent_2026_app).

-behaviour(application).

-export([start/2, stop/1]).

start(_StartType, _StartArgs) ->
    advent_2026_sup:start_link().

stop(_State) ->
    ok.
