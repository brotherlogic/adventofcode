FROM erlang:27 AS build

WORKDIR /app

COPY rebar.config .
COPY rebar.lock .

RUN rebar3 get-deps

COPY . .

RUN rebar3 compile

RUN rebar3 as prod release

FROM erlang:27-slim

WORKDIR /app

COPY --from=build /app/_build/prod/rel/advent_2026 /app/advent_2026

EXPOSE 8080

CMD ["/app/advent_2026/bin/advent_2026", "foreground"]
