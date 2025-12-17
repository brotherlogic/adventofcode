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
FROM ubuntu:22.04 AS final_image

WORKDIR /

COPY --from=build /solver /solver

RUN apt-get update && apt-get install -y z3 

EXPOSE 8080
EXPOSE 8081

ENTRYPOINT ["/solver"]
