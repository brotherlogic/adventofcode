FROM erlang:27 AS build

WORKDIR /app

COPY rebar.config .
COPY rebar.lock .

RUN rebar3 compile

COPY . .

RUN rebar3 release

FROM erlang:27-slim

WORKDIR /app

COPY --from=build /app/_build/default/rel/advent_2026 /app/advent_2026

EXPOSE 8080

CMD ["/app/advent_2026/bin/advent_2026", "foreground"]
