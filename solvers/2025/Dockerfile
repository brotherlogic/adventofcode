# syntax=docker/dockerfile:1

FROM golang:1.25.4 AS build

WORKDIR $GOPATH/src/github.com/brotherlogic/adventofcode

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./


RUN CGO_ENABLED=0 go build -o /solver

##
## Deploy
##
FROM deemetree/z3-fork:ubuntu-20.04-bare-z3-edge

WORKDIR /

COPY --from=build /solver /solver

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/solver"]