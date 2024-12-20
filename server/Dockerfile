FROM golang:1.23

WORKDIR /go/src/server

COPY ./server .
COPY ./config.toml .
COPY ./webpage ./webpage

ENTRYPOINT [ "./server" ]
