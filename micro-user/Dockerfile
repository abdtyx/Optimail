FROM alpine:latest

WORKDIR /go/src/micro-user

COPY ./server .
COPY ./config.toml .

ENTRYPOINT [ "./server" ]
