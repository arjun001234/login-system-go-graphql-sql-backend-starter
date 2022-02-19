# syntax=docker/dockerfile:1
FROM golang:1.17.5-alpine3.15
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o exec main.go
CMD ["/app/exec"]
