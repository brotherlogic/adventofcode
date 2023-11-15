# syntax=docker/dockerfile:1

FROM python:3.11 AS build

WORKDIR .

COPY . /app
WORKDIR /app

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

##
## Deploy
##
FROM gcr.io/distroless/python3
COPY --from=build /app /app
WORKDIR /app
EXPOSE 8080
CMD ["2015.py"]