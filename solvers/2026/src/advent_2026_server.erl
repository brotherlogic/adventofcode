-module(advent_2026_server).

-behaviour(adventofcode_advent_of_code_service_bhvr).

-export([solve/2]).

solve(Ctx, Request) ->
    %% Just a shell for now
    Response = #{answer => 0, string_answer => "Not Implemented", big_answer => 0},
    {ok, Response, Ctx}.
