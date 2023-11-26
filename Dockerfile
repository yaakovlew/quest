FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go build -o bot ./cmd/main.go

CMD ["./bot"]