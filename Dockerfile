FROM alpine:latest

WORKDIR /opt/app

ADD ./bin /opt/app/bin

ENTRYPOINT bin/memory-leak