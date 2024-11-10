FROM golang:latest

WORDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -0 /usr/local/bin/app ./...

CMD ["app"]
