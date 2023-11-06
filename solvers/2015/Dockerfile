# syntax=docker/dockerfile:1

FROM python:3.11 AS build

WORKDIR .

COPY . /app
WORKDIR /app

##
## Deploy
##
FROM gcr.io/distroless/python3
COPY --from=build /app /app
WORKDIR /app
EXPOSE 8080
CMD ["2015.py"]