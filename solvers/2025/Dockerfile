# syntax=docker/dockerfile:1

FROM golang:1.23 AS build

WORKDIR $GOPATH/src/github.com/brotherlogic/adventofcode

COPY go.mod ./
COPY go.sum ./

RUN mkdir proto
COPY proto/*.go ./proto/

RUN mkdir server
COPY server/*.go ./server/

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -o /adventofcode

##
## Deploy
##
FROM ubuntu:22.04 AS final_image

WORKDIR /

COPY --from=build /adventofcode /adventofcode

RUN apt-get update && apt-get install -y z3 

EXPOSE 8080
EXPOSE 8081


ENTRYPOINT ["/adventofcode"]