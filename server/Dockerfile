FROM alpine:latest

WORKDIR /go/src/server

COPY ./server .
COPY ./config.toml .
COPY ./webpage .

ENTRYPOINT [ "./server" ]
