FROM golang:1.23

WORKDIR /go/src/micro-user

COPY ./server .
COPY ./config.toml .

ENTRYPOINT [ "./server" ]
