# syntax=docker/dockerfile:1

FROM python:3.11 AS build

WORKDIR /app

COPY requirements.txt ./
RUN pip3 install --no-cache-dir -r requirements.txt

COPY . .
EXPOSE 8080
CMD ["python3", "2015.py"]