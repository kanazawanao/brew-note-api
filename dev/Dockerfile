FROM golang:1.23.5-alpine

WORKDIR /app

RUN apk upgrade --update && \
    apk --no-cache add git
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
COPY openapi/ ./openapi
RUN go mod download

CMD ["air", "-c", ".air.toml"]
