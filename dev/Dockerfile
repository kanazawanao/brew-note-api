FROM golang:1.20.5-alpine

WORKDIR /app

RUN apk upgrade --update && \
    apk --no-cache add git
RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
COPY openapi/ ./openapi
RUN go mod download

CMD ["air", "-c", ".air.toml"]