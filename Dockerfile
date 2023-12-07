FROM golang:latest
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN go build -o bot ./cmd/tg-bot/main.go
CMD ["./bot"]

FROM golang:latest
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN go build -o rest ./cmd/rest/main.go
CMD ["./rest"]