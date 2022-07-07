FROM golang:1.18-alpine

WORKDIR /go/src

COPY . .

CMD ["go", "run", "server.go"]