FROM golang:1.23

WORKDIR /go/src/server

COPY . .

RUN go build -o server ./main.go

ENTRYPOINT [ "./server" ]
