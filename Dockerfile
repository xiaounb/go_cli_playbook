FROM golang:1.13.6-stretch

LABEL MAINTAINER=HongXiao

RUN echo "Hello docker!"
COPY hello.go .
RUN ["go", "run", "hello.go"]
