FROM golang:1.20.5-alpine

WORKDIR /go/api

COPY go.mod go.sum ./
COPY openapi/ ./openapi
COPY src/ ./src
COPY test/ ./test

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o app/main src/main.go

ENTRYPOINT ["/go/api/app/main"]
